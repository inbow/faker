package server

import (
	"net/http"
)

func (s *Server) Burl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success burl notify")) // nolint:errcheck
}
