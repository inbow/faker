package server

import (
	"encoding/xml"
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
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
	par := &MeetAdsResponse{}
	par.Links = append(par.Links, MeetAdsLink{
		Bid:   s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPV),
		URL:   "http://meetads.com/url",
		Pixel: "http://meetads.com/pixel",
	})

	data, err := xml.MarshalIndent(par, "", "")
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
