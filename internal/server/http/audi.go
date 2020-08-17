package http

import (
	"github.com/oxyd-io/faker/api"
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) AudiPopunder(requestCtx *atreugo.RequestCtx) error {
	response := api.AudiResponse{
		ClickUrl: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Cpc:      generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
