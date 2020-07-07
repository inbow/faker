package server

import (
	"encoding/json"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

// nolint:funlen
func (s *Server) OpenRTBPopunder(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())
	url := string(ctx.QueryArgs().Peek("url"))

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

	ext := struct {
		URL string `json:"url"`
	}{
		URL: s.generator.URLOrDefault(url),
	}

	extBody, _ := json.Marshal(ext)

	bid := openrtb.Bid{
		ID:    randomdata.RandStringRunes(15),
		ImpID: bidRequest.Impressions[0].ID,

		BillingURL: s.generator.OpenRTBURL(generator.BURL),
		NoticeURL:  s.generator.OpenRTBURL(generator.NURL),
		LossURL:    s.generator.OpenRTBURL(generator.LURL),

		Ext: extBody,
	}

	bid.Price = s.generator.PriceOrDefault(price, generator.CPM)

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
