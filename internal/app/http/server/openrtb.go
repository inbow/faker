package server

import (
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

func (s *Server) OpenRTBBanner(ctx *atreugo.RequestCtx) error {
	bidRequest := openrtb.BidRequest{}
	if err := jsoniter.Unmarshal(ctx.PostBody(), &bidRequest); err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	if err := bidRequest.Validate(); err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	if bidRequest.Impressions[0].Banner == nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte("banner object not found"))

		return nil
	}

	bid := openrtb.Bid{
		ID:    randomdata.RandStringRunes(15),
		ImpID: bidRequest.Impressions[0].ID,

		Price:    s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPM),
		AdMarkup: "<html><head></head><body>Hello World</body></html>",

		Width:  bidRequest.Impressions[0].Banner.Width,
		Height: bidRequest.Impressions[0].Banner.Height,

		LossURL:    s.generator.OpenRTBURL(generator.LossURL),
		NoticeURL:  s.generator.OpenRTBURL(generator.NoticeURL),
		BillingURL: s.generator.OpenRTBURL(generator.BiddingURL),
	}

	seatBid := openrtb.SeatBid{
		Bids: []openrtb.Bid{bid},
	}

	bidResponse := openrtb.BidResponse{
		ID:       bidRequest.ID,
		SeatBids: []openrtb.SeatBid{seatBid},
		Currency: "USD",
	}

	data, err := jsoniter.Marshal(bidResponse)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(data)

	return nil
}

func (s *Server) OpenRTBNoticeURL(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success notification url notify"))

	return nil
}

func (s *Server) OpenRTBBiddingURL(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success bidding url notify"))

	return nil
}

func (s *Server) OpenRTBLossURL(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success loss url notify"))

	return nil
}
