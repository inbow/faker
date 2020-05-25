package server

import (
	"encoding/xml"
	"net/http"
	"time"
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

func (s *Server) MeetAdsXML(w http.ResponseWriter, r *http.Request) {
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(r.URL.Query())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			_, _ = w.Write(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return
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
}
