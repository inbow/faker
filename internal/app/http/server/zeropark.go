package server

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	ZeroParkResponse struct {
		XMLName    xml.Name
		Bid        float64 `xml:"bid"`
		URL        string  `xml:"redirecturl"`
		ClickID    string  `xml:"clickid"`
		CampaignID string  `xml:"campaignid"`
	}
)

type (
	ZeroParkPushResponse struct {
		Title         string  `json:"title"`
		Text          string  `json:"description"`
		CPC           float64 `json:"cpc"`
		ClickURL      string  `json:"link"`
		ImageURL      string  `json:"image_url"`
		AdID          string  `json:"ad_id"`
		ImpressionURL string  `json:"imp_url"`

		Token string `json:"token"`
	}
)

func (s *Server) ZeroParkPopunder(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/xml")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	par := ZeroParkResponse{
		XMLName:    xml.Name{Local: "result"},
		Bid:        price,
		URL:        "http://usd.odysseus-nua.com/zcvisitor/c5aa6ef8-b211",
		ClickID:    "c5aa6ef8-b211-11e9-affe-12bbd5c93ce2",
		CampaignID: "c5ad0707-b211-11e9-affe-12bbd5c93ce2",
	}

	data, err := xml.MarshalIndent(par, "", "")
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}

func (s *Server) ZeroParkPush(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	zrppr := ZeroParkPushResponse{
		Title: kozma.Say(),
		Text:  kozma.Say(),

		ClickURL:      "https://push.example.com/c/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImageURL:      "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",
		ImpressionURL: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d/1/{token}",

		CPC: s.generator.PriceOrDefault(price, generator.CPC),

		AdID: uuid.New().String(),

		Token: uuid.New().String(),
	}

	data, err := json.Marshal(zrppr)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
