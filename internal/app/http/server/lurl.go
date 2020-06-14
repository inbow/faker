package server

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
)

func (s *Server) Lurl(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success lurl notify"))

	return nil
}
