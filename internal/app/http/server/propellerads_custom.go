package server

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type (
	PropellerAdsCustomResponse struct {
		Bid float64 `json:"bid"`
		URL string  `json:"url"`
	}
)

func (s *Server) PropellerAdsCustom(w http.ResponseWriter, r *http.Request) {
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(r.URL.Query())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			_, _ = w.Write(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return
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
}
