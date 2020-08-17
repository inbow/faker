package http

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
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

func (s *Server) DatsunNative(requestCtx *atreugo.RequestCtx) error {
	response := datsunResponse{
		Teasers: []Teaser{
			{
				ID:    "20",
				Title: "Play online!",
				Image: "https://api.taptun.com/source/img/c0qxv-5c812253bfc356-00owpx3-4ae",
				Link:  generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
				NURL:  "http://api.taptun.com/v1/feed/nurl?s=c0qxv-5c812253bfc356-00owpx3-4ae",
				CPC:   generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
			},
			{
				ID:    "19",
				Title: "Credit! 0%",
				Image: "https://api.taptun.com/source/img/c0qxv-5c812253c03b52-00w2ndf-40u",
				Link:  generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
				NURL:  "http://api.taptun.com/v1/feed/nurl?s=c0qxv-5c812253c03b52-00w2ndf-40u",
				CPC:   generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
			},
		},
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
