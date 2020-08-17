package http

import (
	"github.com/oxyd-io/faker/internal/generator"
	"github.com/tada-team/kozma"
	"net/http"

	"github.com/savsgio/atreugo/v11"
)

type (
	BuickResponse []*ResponseItem

	ResponseItem struct {
		AdID     string  `json:"id,omitempty"`
		Title    string  `json:"title,omitempty"`
		Text     string  `json:"description,omitempty"`
		CPC      float64 `json:"bid,omitempty"`
		IconURL  string  `json:"icon_url,omitempty"`
		ImageURL string  `json:"image_url,omitempty"`
		ClickURL string  `json:"click_url,omitempty"`
	}
)

func (s *Server) BuickPush(requestCtx *atreugo.RequestCtx) error {
	response := BuickResponse{
		&ResponseItem{
			AdID:     "20",
			Title:    kozma.Say(),
			Text:     kozma.Say(),
			CPC:      generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
			IconURL:  "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
			ImageURL: "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
			ClickURL: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		},
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
