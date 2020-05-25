package server

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	Teaser struct {
		ID    string  `json:"id,omitempty"`
		Title string  `json:"title,omitempty"`
		Image string  `json:"image,omitempty"`
		Link  string  `json:"link,omitempty"`
		NURL  string  `json:"nurl,omitempty"`
		CPC   float64 `json:"cpc,omitempty"`
	}

	datsunResponse struct {
		Teasers []Teaser `json:"teasers"`
	}
)

func (s *Server) Datsun(w http.ResponseWriter, r *http.Request) {
	response := s.NewResponse()
	_, delay, skip := s.RequestValues(r.URL.Query())

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

	par := &datsunResponse{
		Teasers: []Teaser{
			{
				ID:    "20",
				Title: "Play online!",
				Image: "https://api.taptun.com/source/img/c0qxv-5c812253bfc356-00owpx3-4ae",
				Link:  "https://api.taptun.com/v1/click/t?s=c0qxv-5c812253bfc356-00owpx3-4ae",
				NURL:  "http://api.taptun.com/v1/feed/nurl?s=c0qxv-5c812253bfc356-00owpx3-4ae",
				CPC:   s.generator.Price(generator.CPC),
			},
			{
				ID:    "19",
				Title: "Credit! 0%",
				Image: "https://api.taptun.com/source/img/c0qxv-5c812253c03b52-00w2ndf-40u",
				Link:  "https://api.taptun.com/v1/click/t?s=c0qxv-5c812253c03b52-00w2ndf-40u",
				NURL:  "http://api.taptun.com/v1/feed/nurl?s=c0qxv-5c812253c03b52-00w2ndf-40u",
				CPC:   s.generator.Price(generator.CPC),
			},
		},
	}

	data, err := jsoniter.Marshal(par)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data
}
