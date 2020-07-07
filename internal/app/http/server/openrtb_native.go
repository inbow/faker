package server

import (
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"
	"github.com/tidwall/pretty"

	"github.com/oxyd-io/faker/internal/app/generator"
)

// nolint:funlen
func (s *Server) OpenRTBNative(ctx *atreugo.RequestCtx) error {
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
		AdMarkup: openRTBNativeAdm,

		BillingURL: s.generator.OpenRTBURL(generator.BURL),
		NoticeURL:  s.generator.OpenRTBURL(generator.NURL),
		LossURL:    s.generator.OpenRTBURL(generator.LURL),
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

var (
	// nolint:gochecknoglobals
	openRTBNativeAdm = string(pretty.Ugly([]byte(`
{
   "native":{
      "ver":"1.2",
      "link":{
         "url":"https://www.adskeeper.co.uk/"
      },
      "assets":[
         {
            "id":1,
            "required":0,
            "title":{
               "text":"Clean computer in {country}"
            }
         },
         {
            "id":3,
            "required":0,
            "img":{
               "w":360,
               "h":240,
               "type":3,
               "url":"http://main.jpg"
            }
         },
         {
            "id":4,
            "required":0,
            "img":{
               "w":50,
               "h":50,
               "type":1,
               "url":"http://icon.jpg"
            }
         },
         {
            "id":2,
            "required":0,
            "data":{
               "type":1,
               "value":"sponsored by proxeter"
            }
         },
         {
            "id":5,
            "required":0,
            "data":{
               "type":12,
               "value":"Click me if you live in {city}"
            }
         }
      ],
      "imptrackers":[
         "https://notify.adskeeper.co.uk/imp"
      ]
   }
}
`)))
)
