package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"
)

type (
	evadavResponse struct {
		Link  string  `json:"link"`
		Price float64 `json:"cpc"`
	}
)

func (s *Server) Evadav(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/javascript")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	handlerResponse := &evadavResponse{
		Link:  "http://demo.url.link/&demo=1",
		Price: price,
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
