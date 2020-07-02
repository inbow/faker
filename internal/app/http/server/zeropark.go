package server

import (
	"encoding/xml"
	"net/http"

	"github.com/savsgio/atreugo/v11"
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

func (s *Server) ZeroPark(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/xml")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

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

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
