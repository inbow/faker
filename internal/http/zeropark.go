package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/api"
	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) ZeroParkPopunder(requestCtx *atreugo.RequestCtx) error {
	response := api.ZeroparkResponse{
		Cpc:  generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
		Link: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}

func (s *Server) ZeroParkPush(requestCtx *atreugo.RequestCtx) error {
	response := api.ZeroparkResponse{
		Title:       kozma.Say(),
		Description: kozma.Say(),
		Cpc:         generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
		Link:        "https://push.example.com/click/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImageUrl:    "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		AdId:        uuid.New().String(),
		ImpUrl:      "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		Token:       uuid.New().String(),
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
