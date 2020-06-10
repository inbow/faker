package server

import (
	"net/http"
)

func (s *Server) Lurl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success lurl notify")) // nolint:errcheck
}
