package http

import (
	"net/http"

	"github.com/oxyd-io/faker/api"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) PeugeotPush(requestCtx *atreugo.RequestCtx) error {
	peugeotResponse := api.PeugeotResponse{}

	peugeotResponseItem := &api.PeugeotItemResponse{
		Cpc:         generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
		Url:         generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Title:       kozma.Say(),
		Description: kozma.Say(),
		Icon:        "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Image:       "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		AdId:        "645658476",
	}

	peugeotResponse.Result = append(peugeotResponse.Result, peugeotResponseItem)

	return requestCtx.JSONResponse(&peugeotResponse, http.StatusOK)
}

func (s *Server) PeugeotPopunder(requestCtx *atreugo.RequestCtx) error {
	peugeotResponse := api.PeugeotResponse{}

	peugeotResponseItem := &api.PeugeotItemResponse{
		Url: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Cpc: generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
	}

	peugeotResponse.Result = append(peugeotResponse.Result, peugeotResponseItem)

	return requestCtx.JSONResponse(&peugeotResponse, http.StatusOK)
}
