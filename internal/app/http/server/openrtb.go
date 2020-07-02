package server

import (
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

// nolint:funlen
func (s *Server) OpenRTB(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	bidRequest := openrtb.BidRequest{}
	if err := jsoniter.Unmarshal(ctx.PostBody(), &bidRequest); err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	if err := bidRequest.Validate(); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = []byte(err.Error())

		return nil
	}

	bid := openrtb.Bid{
		ID:       randomdata.RandStringRunes(15),
		ImpID:    bidRequest.Impressions[0].ID,
		AdMarkup: s.generator.AdMarkup(),

		Width:  bidRequest.Impressions[0].Banner.Width,
		Height: bidRequest.Impressions[0].Banner.Height,

		BillingURL: s.generator.URL(generator.BURL),
		NoticeURL:  s.generator.URL(generator.NURL),
		LossURL:    s.generator.URL(generator.LURL),
	}

	if price != 0 {
		bid.Price = price
	} else {
		bid.Price = s.generator.Price(generator.CPM)
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
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}

func (s *Server) NotificationURL(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success nurl notify"))

	return nil
}

func (s *Server) BiddingURL(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success burl notify"))

	return nil
}

func (s *Server) LossURL(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("Success lurl notify"))

	return nil
}
