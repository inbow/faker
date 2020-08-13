package http

import (
	"encoding/json"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) OpenRTBPopunder(requestCtx *atreugo.RequestCtx) error {
	bidRequest := openrtb.BidRequest{}
	if err := jsoniter.Unmarshal(requestCtx.PostBody(), &bidRequest); err != nil {
		requestCtx.SetStatusCode(http.StatusBadGateway)
		requestCtx.SetBody([]byte(err.Error()))

		return nil
	}

	if err := bidRequest.Validate(); err != nil {
		requestCtx.SetStatusCode(http.StatusBadGateway)
		requestCtx.SetBody([]byte(err.Error()))

		return nil
	}

	ext := struct {
		URL string `json:"url"`
	}{
		URL: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
	}

	extBody, _ := json.Marshal(ext)

	bid := openrtb.Bid{
		ID:    randomdata.RandStringRunes(15),
		ImpID: bidRequest.Impressions[0].ID,

		Price: generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPM),

		LossURL:    generator.OpenRTBURL(s.config, generator.LossURL),
		NoticeURL:  generator.OpenRTBURL(s.config, generator.NoticeURL),
		BillingURL: generator.OpenRTBURL(s.config, generator.BiddingURL),

		Ext: extBody,
	}

	seatBid := openrtb.SeatBid{
		Bids: []openrtb.Bid{bid},
	}

	bidResponse := openrtb.BidResponse{
		ID:       bidRequest.ID,
		SeatBids: []openrtb.SeatBid{seatBid},
		Currency: "USD",
	}

	return requestCtx.JSONResponse(&bidResponse, http.StatusOK)
}
