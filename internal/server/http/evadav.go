package http

import (
	"github.com/google/uuid"
	"github.com/tada-team/kozma"
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/api"
	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) EvadavPopunder(requestCtx *atreugo.RequestCtx) error {
	handlerResponse := api.EvadavPopunderResponse{
		Link: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Cpc:  generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
	}

	return requestCtx.JSONResponse(&handlerResponse, http.StatusOK)
}

func (s *Server) EvadavPush(requestCtx *atreugo.RequestCtx) error {
	aesearchResponse := api.EvadavPushResponse{
		Id:    uuid.New().String(),
		Title: kozma.Say(),
		Descr: kozma.Say(),
		Icon:  "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Image: "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Link:  generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Cpc:   generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
	}

	return requestCtx.JSONResponse(&aesearchResponse, http.StatusOK)
}
