package server

import (
	"encoding/xml"
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	IntangoResponse struct {
		XMLName xml.Name        `xml:"results"`
		Result  []IntangoResult `xml:"result"`
	}

	IntangoResult struct {
		Title       string  `xml:"title"`
		Description string  `xml:"description"`
		DisplayURL  string  `xml:"displayurl"`
		URL         string  `xml:"url"`
		DID         string  `xml:"did"`
		Bid         float64 `xml:"bid"`
	}
)

func (s *Server) IntangoXML(ctx *atreugo.RequestCtx) error {
	iar := &IntangoResponse{}
	iar.Result = append(iar.Result, IntangoResult{
		Title:       "PopAd Title",
		Description: "PopAd Desc1. PopAd Desc2",
		DisplayURL:  "PopAds.com",
		URL:         "http://mybestdc.com/aS/feedclick",
		DID:         "2fa9106c-2a26-4d1e-9ffe-fcea038b402a",
		Bid:         s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPV),
	})

	data, err := xml.MarshalIndent(iar, "", "")
	if err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	ctx.Response.Header.Set("Content-Type", "application/xml")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(data)

	return nil
}
