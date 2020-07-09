package server

import (
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	Data struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DisplayURL  string `json:"displayURL"`
		Redirect    string `json:"redirect"`
		CPC         string `json:"cpc"`
		Image       string `json:"image"`
	}

	VolvoResponse struct {
		Data []Data `json:"data,omitempty"`

		Errors []string `json:"errors,omitempty"`
	}
)

func (s *Server) Volvo(ctx *atreugo.RequestCtx) error {
	price := s.generator.PriceOrDefault(ctx.UserValue(string(Price)).(float64), generator.CPC)

	handlerResponse := &VolvoResponse{
		Data: []Data{
			{
				Title:       "Yukon Gold Casino",
				Description: "125 Chances to Win Jackpot. One spin can change your life!",
				DisplayURL:  "www.yukongoldcasino.eu",
				Redirect:    "https://rsalchk.com/performance/bdv_rd.dbm?enparms2=qwerty",
				CPC:         strconv.FormatFloat(price, 'g', 6, 64),
				Image:       "//bvadimgs.scdn7.secure.raxcdn.com/bidvertiser/images/ad_image/9/96226.png",
			},
		},
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.SetBody([]byte(err.Error()))

		return nil
	}

	ctx.Response.Header.Set("Content-Type", "application/javascript")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(data)

	return nil
}
