package http

import (
	"encoding/xml"
	"net/http"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/generator"
)

type (
	ZeroParkPopunderResponse struct {
		XMLName xml.Name

		Bid        float64 `xml:"bid"`
		URL        string  `xml:"redirecturl"`
		ClickID    string  `xml:"clickid"`
		CampaignID string  `xml:"campaignid"`
	}

	ZeroParkPushResponse struct {
		Title         string  `json:"title"`
		Text          string  `json:"description"`
		CPC           float64 `json:"cpc"`
		ClickURL      string  `json:"link"`
		ImageURL      string  `json:"image_url"`
		AdID          string  `json:"ad_id"`
		ImpressionURL string  `json:"imp_url"`
		Token         string  `json:"token"`
	}
)

func (s *Server) ZeroParkPopunder(requestCtx *atreugo.RequestCtx) error {
	response := ZeroParkPopunderResponse{
		XMLName:    xml.Name{Local: "result"},
		Bid:        generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
		URL:        generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		ClickID:    "c5aa6ef8-b211-11e9-affe-12bbd5c93ce2",
		CampaignID: "c5ad0707-b211-11e9-affe-12bbd5c93ce2",
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}

func (s *Server) ZeroParkPush(requestCtx *atreugo.RequestCtx) error {
	response := ZeroParkPushResponse{
		Title: kozma.Say(),
		Text:  kozma.Say(),

		ClickURL:      "https://push.example.com/c/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImageURL:      "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImpressionURL: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",

		CPC: generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),

		AdID: uuid.New().String(),

		Token: uuid.New().String(),
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
