package http

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/api"
	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) MackPush(requestCtx *atreugo.RequestCtx) error {
	mackResponse := api.MackPushResponse{
		Id:          "4687654",
		Crid:        "153468",
		Title:       kozma.Say(),
		Description: kozma.Say(),
		Link:        generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Icon:        "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Image:       "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Nurl:        "https://push.example.com/win/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Price:       generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
	}

	return requestCtx.JSONResponse(&mackResponse, http.StatusOK)
}
