package http

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"

	"github.com/oxyd-io/faker/internal/generator"
)

type (
	PropellerAdsPopunderResponse struct {
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

func (s *Server) PropellerAdsPopunder(requestCtx *atreugo.RequestCtx) error {
	response := PropellerAdsPopunderResponse{
		Bid: generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPV),
		URL: generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}

func (s *Server) PropellerAdsPush(requestCtx *atreugo.RequestCtx) error {
	response := PropellerAdsPushResponse{
		Ads: Ads{
			BannerID:      123213,
			CampaignID:    239,
			Title:         "CONGRATS TO THE WINNERS!",
			Text:          "Tap now to see if you won $1,909,349!",
			Icon:          "https://offerimage.com/www/images/icon.jpg",
			Image:         "https://offerimage.com/www/images/img.jpg",
			ClickURL:      generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
			ImpressionURL: "https://offers.propellerads.com/some_log_impression_path?some_param=1",
			CPCRate:       generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
			Rate:          generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
			RateModel:     "cpc",
		},
	}

	return requestCtx.JSONResponse(&response, http.StatusOK)
}
