package http

import (
	"encoding/xml"
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

type (
	MeetAdsResponse struct {
		XMLName xml.Name `xml:"result"`

		Links []MeetAdsLink `xml:"link"`
	}

	MeetAdsLink struct {
		Bid   float64 `xml:"bid,attr"`
		URL   string  `xml:"url,attr"`
		Pixel string  `xml:"pixel,attr"`
	}
)

func (s *Server) MeetAds(requestCtx *atreugo.RequestCtx) error {
	par := &MeetAdsResponse{}
	par.Links = append(par.Links, MeetAdsLink{
		Bid:   generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
		URL:   generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Pixel: "http://meetads.com/pixel",
	})

	data, err := xml.MarshalIndent(par, "", "")
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
