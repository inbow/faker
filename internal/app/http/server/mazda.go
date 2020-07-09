package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	mazdaResponse struct {
		Link  string  `json:"url"`
		Price float64 `json:"cpc"`
	}
)

func (s *Server) Mazda(ctx *atreugo.RequestCtx) error {
	handlerResponse := &mazdaResponse{
		Link:  "http://demo.url.link/&demo=1",
		Price: s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPV),
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	ctx.Response.Header.Set("Content-Type", "application/javascript")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(data)

	return nil
}
