package server

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"
)

type (
	PropellerAdsCustomResponse struct {
		Bid float64 `json:"bid"`
		URL string  `json:"url"`
	}
)

func (s *Server) PropellerAdsCustom(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(ctx.QueryArgs())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return nil
	}

	par := PropellerAdsCustomResponse{
		Bid: price,
		URL: "http://digitaldsp.com/api/win_request?p=Z",
	}

	data, err := jsoniter.Marshal(par)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
