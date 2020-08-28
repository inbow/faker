package http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/api"
	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) ChevroletPush(requestCtx *atreugo.RequestCtx) error {
	chevroletResponse := api.ChevroletPushResponse{
		ImpressionKey:       uuid.New().String(),
		ImpressionServedUrl: fmt.Sprintf("https://%v/api/v1/chevrolet/impression", s.config.HTTP.Domain),
		SlotCount:           1,
	}

	chevroletResponseItem := &api.ChevroletPushItemResponse{
		Slot:  1,
		Url:   generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Title: kozma.Say(),
		Image: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Cpc:   generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
	}

	chevroletResponse.Articles = append(chevroletResponse.Articles, chevroletResponseItem)

	return requestCtx.JSONResponse(&chevroletResponse, http.StatusOK)
}

func (s *Server) ChevroletImpression(requestCtx *atreugo.RequestCtx) error {
	requestCtx.SetStatusCode(http.StatusOK)
	requestCtx.SetBody([]byte("Success impression notification by POST"))

	return nil
}
