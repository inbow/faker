package http

import (
	"net/http"

	"github.com/oxyd-io/faker/api"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) FordPush(requestCtx *atreugo.RequestCtx) error {
	aesearchResponse := api.FordPushResponse{
		Status: "success",
		Data: &api.FordPushItemResponse{
			Cpc:         generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
			Title:       kozma.Say(),
			Description: kozma.Say(),
			Icon:        "https://push.example.com//dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
			Image:       "https://push.example.com/imagicone/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
			Link:        generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		},
	}

	return requestCtx.JSONResponse(&aesearchResponse, http.StatusOK)
}
