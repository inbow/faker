package http

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

type (
	mazdaResponse struct {
		Link  string  `json:"url"`
		Price float64 `json:"cpc"`
	}
)

func (s *Server) Mazda(requestCtx *atreugo.RequestCtx) error {
	response := mazdaResponse{
		Link:  generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Price: generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
