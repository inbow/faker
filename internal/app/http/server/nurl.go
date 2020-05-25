package server

import (
	"net/http"
)

func (s *Server) Nurl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success nurl notify")) // nolint:errcheck
}
