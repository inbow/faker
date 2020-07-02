package server

import (
	"encoding/xml"
	"net/http"

	"github.com/savsgio/atreugo/v11"
)

type (
	MeetAdsResponse struct {
		XMLName xml.Name      `xml:"result"`
		Links   []MeetAdsLink `xml:"link"`
	}

	MeetAdsLink struct {
		Bid   float64 `xml:"bid,attr"`
		URL   string  `xml:"url,attr"`
		Pixel string  `xml:"pixel,attr"`
	}
)

func (s *Server) MeetAdsXML(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/xml")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	par := &MeetAdsResponse{}
	par.Links = append(par.Links, MeetAdsLink{
		Bid:   price,
		URL:   "http://meetads.com/url",
		Pixel: "http://meetads.com/pixel",
	})

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
