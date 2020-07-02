package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/savsgio/atreugo/v11"
)

type (
	PropellerAdsCustomResponse struct {
		Bid float64 `json:"bid"`
		URL string  `json:"url"`
	}
)

type (
	Ads struct {
		BannerID      int     `json:"banner_id,omitempty"`
		CampaignID    int     `json:"campaign_id,omitempty"`
		Title         string  `json:"title,omitempty"`
		Text          string  `json:"text,omitempty"`
		Icon          string  `json:"icon,omitempty"`
		Image         string  `json:"image,omitempty"`
		ClickURL      string  `json:"click_url,omitempty"`
		ImpressionURL string  `json:"impression_url,omitempty"`
		CPCRate       float64 `json:"cpc_rate,omitempty"`
		Rate          float64 `json:"rate,omitempty"`
		RateModel     string  `json:"rate_model,omitempty"`

		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}

	PropellerAdsPushResponse struct {
		Ads Ads `json:"ads"`
	}
)

func (s *Server) PropellerAdsCustom(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	par := PropellerAdsCustomResponse{
		Bid: price,
		URL: "http://digitaldsp.com/api/win_request?p=Z",
	}

	data, err := jsoniter.Marshal(par)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}

func (s *Server) PropellerAdsPush(ctx *atreugo.RequestCtx) error {
	response := s.NewResponse()
	price := s.price(ctx.QueryArgs())

	defer func() {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			ctx.SetBody(response.Body)
		}
	}()

	par := PropellerAdsPushResponse{
		Ads: Ads{
			BannerID:      123213,
			CampaignID:    239,
			Title:         "CONGRATS TO THE WINNERS!",
			Text:          "Tap now to see if you won $1,909,349!",
			Icon:          "https://offerimage.com/www/images/icon.jpg",
			Image:         "https://offerimage.com/www/images/img.jpg",
			ClickURL:      "https://offers.propellerads.com/some_log_click_path?some_param=1",
			ImpressionURL: "https://offers.propellerads.com/some_log_impression_path?some_param=1",
			CPCRate:       price,
			Rate:          price,
			RateModel:     "cpc",
		},
	}

	data, err := jsoniter.Marshal(par)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())

		return nil
	}

	response.StatusCode = http.StatusOK
	response.Body = data

	return nil
}
