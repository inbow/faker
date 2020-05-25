package server

import (
	"encoding/xml"
	"net/http"
	"time"
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

func (s *Server) IntangoXML(w http.ResponseWriter, r *http.Request) {
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
}
