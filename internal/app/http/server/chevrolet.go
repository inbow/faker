package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/oxyd-io/atom"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/app/generator"
)

func (s *Server) ChevroletPush(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	chevroletResponse := atom.ChevroletPushResponse{
		ImpressionKey:       uuid.New().String(),
		ImpressionServedUrl: fmt.Sprintf("https://%v:%v/api/v1/chevrolet/impression", s.config.HTTP.Host, s.config.HTTP.Port),
		SlotCount:           1,
	}

	chevroletResponseItem := &atom.ChevroletPushItemResponse{
		Slot:  1,
		Url:   "https://push.example.com/c/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Title: kozma.Say(),
		Image: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Cpc:   s.generator.PriceOrDefault(price, generator.CPC),
	}

	chevroletResponse.Articles = append(chevroletResponse.Articles, chevroletResponseItem)

	data, err := json.Marshal(&chevroletResponse)
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
