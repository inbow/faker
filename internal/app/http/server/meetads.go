package server

import (
	"encoding/xml"
	"net/http"
	"time"

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
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
