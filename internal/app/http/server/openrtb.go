package server

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bsm/openrtb"
	jsoniter "github.com/json-iterator/go"

	"github.com/oxyd-io/faker/internal/app/generator"
)

// nolint:funlen
func (s *Server) OpenRTB(w http.ResponseWriter, r *http.Request) {
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(r.URL.Query())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			_, _ = w.Write(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = []byte(err.Error())

		return
	}

	bidRequest := openrtb.BidRequest{}
	if err := jsoniter.Unmarshal(data, &bidRequest); err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return
	}

	if err := bidRequest.Validate(); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = []byte(err.Error())

		return
	}

	// nurlUrl := "http://" + s.config.HTTP.Host + ":" + strconv.Itoa(s.config.HTTP.Port) + "/nurl"
	// if len(query.Get("nurlUrl")) > 0 {
	// 	nurlUrl = query.Get("nurlUrl")
	// }
	//
	// burlUrl := "http://" + s.config.HTTP.Host + ":" + strconv.Itoa(s.config.HTTP.Port) + "/burl"
	// if len(query.Get("burlUrl")) > 0 {
	// 	burlUrl = query.Get("burlUrl")
	// }

	bid := openrtb.Bid{
		ID:       randomdata.RandStringRunes(15),
		ImpID:    bidRequest.Imp[0].ID,
		AdMarkup: s.generator.AdMarkup(),
		BURL:     s.generator.URL(generator.BURL),
		NURL:     s.generator.URL(generator.NURL),
		LURL:     s.generator.URL(generator.LURL),

		W: bidRequest.Imp[0].Banner.W,
		H: bidRequest.Imp[0].Banner.H,

		// NURL: nurlUrl,
		// BURL: burlUrl,
	}

	if price != 0 {
		bid.Price = price
	} else {
		bid.Price = s.generator.Price(generator.CPM)
	}

	seatBid := openrtb.SeatBid{
		Bid: []openrtb.Bid{bid},
	}

	bidResponse := openrtb.BidResponse{
		ID:       bidRequest.ID,
		SeatBid:  []openrtb.SeatBid{seatBid},
		Currency: "USD",
	}

	data, err = jsoniter.Marshal(bidResponse)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data
}
