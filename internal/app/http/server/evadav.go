package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/oxyd-io/atom"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

func (s *Server) EvadavPopunder(ctx *atreugo.RequestCtx) error {
	handlerResponse := &atom.EvadavPopunderResponse{
		Link: s.generator.URLOrDefault(ctx.UserValue(string(URL)).(string)),
		Cpc:  s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPV),
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
