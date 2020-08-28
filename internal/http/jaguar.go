package http

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/api"
	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) JaguarPush(requestCtx *atreugo.RequestCtx) error {
	jaguarResponse := api.JaguarResponse{}

	jaguarResponseItem := &api.JaguarItemResponse{
		Title:         kozma.Say(),
		Text:          kozma.Say(),
		Icon:          "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Image:         "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		ClickUrl:      generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		ImpressionUrl: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		CpcRate:       generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
		CampaignId:    5648,
	}

	jaguarResponse.Articles = append(jaguarResponse.Articles, jaguarResponseItem)

	return requestCtx.JSONResponse(&jaguarResponse, http.StatusOK)
}

func (s *Server) JaguarNative(requestCtx *atreugo.RequestCtx) error {
	jaguarResponse := api.JaguarResponse{}

	jaguarResponseItem := &api.JaguarItemResponse{
		Title:         kozma.Say(),
		Text:          kozma.Say(),
		Icon:          "https://example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Image:         "https://example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		ClickUrl:      generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		ImpressionUrl: "https://example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		CpcRate:       generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
		CampaignId:    468574,
	}

	jaguarResponse.Articles = append(jaguarResponse.Articles, jaguarResponseItem)

	return requestCtx.JSONResponse(&jaguarResponse, http.StatusOK)
}
