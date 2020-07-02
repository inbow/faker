package server

import (
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"
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
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/javascript")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

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
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
