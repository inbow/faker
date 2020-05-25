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
func (s *Server) OpenRTBVast(w http.ResponseWriter, r *http.Request) {
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
		AdMarkup: VastADM,
		BURL:     s.generator.URL(generator.BURL),
		NURL:     s.generator.URL(generator.NURL),
		LURL:     s.generator.URL(generator.LURL),

		// W: bidRequest.Imp[0].Banner.W,
		// H: bidRequest.Imp[0].Banner.H,
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

const (
	VastADM = `<VAST version=\"3.0\"><Ad><InLine><AdSystem version=\"3.0\"><![CDATA[iabtechlab]]></AdSystem><AdTitle><![CDATA[Erodate.pl]]></AdTitle><Impression><![CDATA[http://spilaggroved.xyz/dambFczddeGfVg2hZiWj5k0lPm2nFokpSqWrQs9tMujvgw3xOyDzMA1BJCmDVE2FZGWH5I0JVKHLlMwNZOTP0Q1RJSnTBUhVeWWX9Y1ZdaDb0c4dMeDfAgwhMiDjAkmlcmmnVo2pZqWr5s1tZuTv0wxxMyDzAAwBMCDDAEwFJGnHNIJJZKDL0M0NOODPYQ0RMSzTIU2VNWTXAY1ZJanbNcJdZeDfIg9hJinjNkJlZmDnMo9pJqnrRsvtau2vVwuxPySzZA6BbC2D5ElFSGWHQI9JNKDLAM4NNOTPAQ3RNSwT]]></Impression><Creatives><Creative><Linear><Duration>00:00:05</Duration><Icons></Icons><TrackingEvents><Tracking event=\"start\"><![CDATA[http://spilaggroved.xyz/dlmmFnzodpGqVr2sZtWu5v0wPx2yFzkASBWCQD9EMFjGgH3IOJDKML1MJNmOVP2QZRWS5T0UVVHWlXwYZZTa0bycJdneBfhgehWi9j1kdlDm0n4oMpDqArwsMtDuAvmwcxmyVz2AZBWC5D1EZFTG0HxIMJDKALwMMNDOAPwQJRnSNTJUZVDW0X0YOZDaYb0cMdzeIf2gNhTiAj1kJlnmNnJoZpDqIr9sJtnuNvJwZxDyMz9AJBnCRDvEaF2GVHuIPJSKZL6MbN2O5PlQSRWSQT9UNVDWAX4YNZTaAb3cNdwe]]></Tracking></TrackingEvents><VideoClicks><ClickThrough><![CDATA[http://spilaggroved.xyz/dbmcFdzedfGgVh2iZjWk5l0mPn2oFpkqSrWsQt9uMvjwgx3yOzDAMB1CJDmEVF2GZHWI5J0KVLHMlNwOZPTQ0RzSJTnUBVhWeXWY9Z1adbDc0d4eMfDgAhwiMjDkAlmmcnmoVpkqarXsJtluYv3wRxVyczmAwB9CaDHERF0GcHHIMJlKML0MENlOMPkQYRlSMTkUZV3WdX3YcZuaZbXcJdveZfGgFh0iZjSk5lwmbnCoUpyqRrnsJtluZv2wlxzydzHAJBhCdDGElFvGbHiIUJyKRLnMJNvOdPGQFR0SZTSUUVyWRXkYFZXabbFcNdyeYf1gYhyiYj3kdl5mSnFoFpJqcrUs1truZvWwpxayJzTANBGCcDyEUF1GQHnINJvKdLXMJNjOZPUQlRkSJTTUVVEWJXTYNZEaMbzcldheYfmgQhwiYjTkAl5mMnWoNplqNrjsUtxuMvTwJxkyZzjAhBjCNDWEFFlGOHDIkJwKYLzMZNhOZPGQMRlSMTjUZVzWJXTYVZCacb3cVdieSfWgQhliNjUkQllmMn0oRpEqRrUsVtJuSvEwcxyyOzEAgBmCcDmEVF2GZHWI5J1KZLTM0NxOMPDQARwSMTDUAVwWJXnYNZJaZbDc0d0eOfDgYh0iMjzkIl2mNnToAp1qJrnsNtJuZvDwIx9yJznANBJCZDDEMF9GJHnIRJvKaL2MVNuOPPSQZR6SbT2U5VlWSXWYQZ9aNbDcAd4eNfTgAh3iNjwk]]></ClickThrough></VideoClicks><MediaFiles><MediaFile delivery=\"progressive\" type=\"video/mp4\" codec=\"h264\" bitrate=\"465641\" width=\"560\" height=\"320\"><![CDATA[//10-81.s.cdn15.com/cr/28/118933/287835_a7ef3.mp4]]></MediaFile><MediaFile delivery=\"progressive\" type=\"video/webm\" codec=\"vp8\" width=\"560\" height=\"320\"><![CDATA[//10-81.s.cdn15.com/cr/28/118933/287835_a7ef31.webm]]></MediaFile><MediaFile delivery=\"progressive\" type=\"video/flv\" codec=\"flv1\" bitrate=\"466000\" width=\"560\" height=\"320\"><![CDATA[//10-81.s.cdn15.com/cr/28/118933/287835_a7ef32.flv]]></MediaFile></MediaFiles></Linear></Creative></Creatives><Description></Description><Survey></Survey></InLine></Ad></VAST>`
)
