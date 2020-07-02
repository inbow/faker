package server

// Response for ContentAd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	ChevroletPushResponse struct {
		ImpressionKey string         `json:"impression_key"`
		ImpressionURL string         `json:"impression_served_url"`
		SlotCount     int            `json:"slot_count"`
		ResponseItem  []ResponseItem `json:"articles"`
	}

	ResponseItem struct {
		Slot        int     `json:"slot"`
		ClickURL    string  `json:"url"`
		Title       string  `json:"title"`
		ImageURL    string  `json:"image"`
		Sponsored   bool    `json:"sponsored"`
		SponsoredBy string  `json:"sponsored_by"`
		CPC         float64 `json:"cpc"`
		Score       float64 `json:"score"`
	}
)

func (s *Server) ChevroletPush(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	_ = s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	chevroletResponseItem := ResponseItem{
		Slot:     1,
		ClickURL: "https://push.example.com/c/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Title:    kozma.Say(),
		ImageURL: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		CPC:      s.generator.Price(generator.CPC),
	}

	chevroletResponse := &ChevroletPushResponse{
		ImpressionKey: uuid.New().String(),
		ImpressionURL: fmt.Sprintf("https://%v:%v/api/v1/chevrolet/impression", s.config.HTTP.Host, s.config.HTTP.Port),
		SlotCount:     1,
	}

	chevroletResponse.ResponseItem = append(chevroletResponse.ResponseItem, chevroletResponseItem)

	data, err := json.Marshal(chevroletResponse)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}

func (s *Server) ChevroletImpression(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success post impression notification"))

	return nil
}
