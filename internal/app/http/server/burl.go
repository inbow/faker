package server

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
)

func (s *Server) Burl(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success burl notify"))

	return nil
}
