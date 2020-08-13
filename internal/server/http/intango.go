package http

import (
	"encoding/xml"
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

type (
	IntangoResponse struct {
		XMLName xml.Name `xml:"results"`

		Result []IntangoResult `xml:"result"`
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

func (s *Server) Intango(requestCtx *atreugo.RequestCtx) error {
	iar := &IntangoResponse{}
	iar.Result = append(iar.Result, IntangoResult{
		Title:       "PopAd Title",
		Description: "PopAd Desc1. PopAd Desc2",
		DisplayURL:  "PopAds.com",
		URL:         generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		DID:         "2fa9106c-2a26-4d1e-9ffe-fcea038b402a",
		Bid:         generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
	})

	data, err := xml.MarshalIndent(iar, "", "")
	if err != nil {
		requestCtx.SetStatusCode(http.StatusBadGateway)
		requestCtx.SetBody([]byte(err.Error()))

		return nil
	}

	requestCtx.Response.Header.Set("Content-Type", "application/xml")
	requestCtx.SetStatusCode(http.StatusOK)
	requestCtx.SetBody(data)

	return nil
}
