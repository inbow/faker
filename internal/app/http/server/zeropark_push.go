package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	ZeroparkPushResponse struct {
		Title         string  `json:"title"`
		Text          string  `json:"description"`
		CPC           float64 `json:"cpc"`
		ClickURL      string  `json:"link"`
		ImageURL      string  `json:"image_url"`
		AdID          string  `json:"ad_id"`
		ImpressionURL string  `json:"imp_url"`

		Token string `json:"token"`
	}
)

func (s *Server) ZeroparkPush(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	_ = s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	zrppr := ZeroparkPushResponse{
		Title: kozma.Say(),
		Text:  kozma.Say(),

		ClickURL:      "https://push.example.com/c/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImageURL:      "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImpressionURL: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",

		CPC: s.generator.Price(generator.CPC),

		AdID: uuid.New().String(),

		Token: uuid.New().String(),
	}

	data, err := json.Marshal(zrppr)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
