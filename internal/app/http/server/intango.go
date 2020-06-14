package server

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/savsgio/atreugo/v11"
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
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(ctx.QueryArgs())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		ctx.Response.Header.Set("Content-Type", "application/xml")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return nil
	}

	iar := &IntangoResponse{}
	iar.Result = append(iar.Result, IntangoResult{
		Title:       "PopAd Title",
		Description: "PopAd Desc1. PopAd Desc2",
		DisplayURL:  "PopAds.com",
		URL:         "http://mybestdc.com/aS/feedclick",
		DID:         "2fa9106c-2a26-4d1e-9ffe-fcea038b402a",
		Bid:         price,
	})

	data, err := xml.MarshalIndent(iar, "", "")
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
