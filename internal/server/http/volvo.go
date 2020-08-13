package http

import (
	"net/http"
	"strconv"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
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

func (s *Server) Volvo(requestCtx *atreugo.RequestCtx) error {
	price := generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC)

	response := VolvoResponse{
		Data: []Data{
			{
				Title:       "Yukon Gold Casino",
				Description: "125 Chances to Win Jackpot. One spin can change your life!",
				DisplayURL:  "www.yukongoldcasino.eu",
				Redirect:    generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
				CPC:         strconv.FormatFloat(price, 'g', 6, 64),
				Image:       "//bvadimgs.scdn7.secure.raxcdn.com/bidvertiser/images/ad_image/9/96226.png",
			},
		},
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
