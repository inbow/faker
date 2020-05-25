package server

import (
	"encoding/xml"
	"net/http"
	"time"
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

func (s *Server) ZeroPark(w http.ResponseWriter, r *http.Request) {
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
	}

	response.StatusCode = http.StatusOK
	response.Body = data
}
