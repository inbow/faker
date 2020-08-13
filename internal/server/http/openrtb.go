package http

import (
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) OpenRTBBanner(requestCtx *atreugo.RequestCtx) error {
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

	if bidRequest.Impressions[0].Banner == nil {
		requestCtx.SetStatusCode(http.StatusBadGateway)
		requestCtx.SetBody([]byte("banner object not found"))

		return nil
	}

	bid := openrtb.Bid{
		ID:    randomdata.RandStringRunes(15),
		ImpID: bidRequest.Impressions[0].ID,

		Price:    generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPM),
		AdMarkup: "<html><head></head><body>Hello World</body></html>",

		Width:  bidRequest.Impressions[0].Banner.Width,
		Height: bidRequest.Impressions[0].Banner.Height,

		LossURL:    generator.OpenRTBURL(s.config, generator.LossURL),
		NoticeURL:  generator.OpenRTBURL(s.config, generator.NoticeURL),
		BillingURL: generator.OpenRTBURL(s.config, generator.BiddingURL),
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

func (s *Server) OpenRTBNoticeURL(requestCtx *atreugo.RequestCtx) error {
	requestCtx.SetStatusCode(http.StatusOK)
	requestCtx.SetBody([]byte("Success notification url notify"))

	return nil
}

func (s *Server) OpenRTBBiddingURL(requestCtx *atreugo.RequestCtx) error {
	requestCtx.SetStatusCode(http.StatusOK)
	requestCtx.SetBody([]byte("Success bidding url notify"))

	return nil
}

func (s *Server) OpenRTBLossURL(requestCtx *atreugo.RequestCtx) error {
	requestCtx.SetStatusCode(http.StatusOK)
	requestCtx.SetBody([]byte("Success loss url notify"))

	return nil
}
