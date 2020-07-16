package server

import (
	"github.com/tada-team/kozma"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/oxyd-io/atom"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

func (s *Server) FordPush(ctx *atreugo.RequestCtx) error {
	handlerResponse := &atom.FordPushResponse{
		Status: "s",
		Data: &atom.FordPushItemResponse{
			Cpc:         s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPC),
			Title:       kozma.Say(),
			Description: kozma.Say(),
			Icon:        "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-ford",
			Image:       "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-ford",
			Link:        s.generator.URLOrDefault(ctx.UserValue(string(URL)).(string)),
		},
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(data)

	return nil
}
