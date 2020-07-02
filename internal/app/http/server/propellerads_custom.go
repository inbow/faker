package server

import (
	"net/http"

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
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	par := PropellerAdsCustomResponse{
		Bid: price,
		URL: "http://digitaldsp.com/api/win_request?p=Z",
	}

	data, err := jsoniter.Marshal(par)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
