package server

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type (
	evadavResponse struct {
		Link  string  `json:"link"`
		Price float64 `json:"cpc"`
	}
)

func (s *Server) Evadav(w http.ResponseWriter, r *http.Request) {
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(r.URL.Query())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			_, _ = w.Write(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return
	}

	handlerResponse := &evadavResponse{
		Link:  "http://demo.url.link/&demo=1",
		Price: price,
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data
}
