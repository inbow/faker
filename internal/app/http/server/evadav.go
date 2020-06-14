package server

import (
	"net/http"
	"time"

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
	price, delay, skip := s.RequestValues(ctx.QueryArgs())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		ctx.Response.Header.Set("Content-Type", "application/javascript")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return nil
	}

	handlerResponse := &evadavResponse{
		Link:  "http://demo.url.link/&demo=1",
		Price: price,
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
