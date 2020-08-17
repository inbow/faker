package http

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
	"github.com/tada-team/kozma"

	"github.com/oxyd-io/faker/api"
	"github.com/oxyd-io/faker/internal/generator"
)

func (s *Server) AesearchPush(requestCtx *atreugo.RequestCtx) error {
	aesearchResponse := api.AesearchResponse{}

	aesearchResponseItem := &api.AesearchItemResponse{
		Title: kozma.Say(),
		Descr: kozma.Say(),
		Icon:  "https://push.example.com/icon/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Url:   generator.URLOrDefault(requestCtx.UserValue(string(URL)).(string)),
		Image: "https://push.example.com/image/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Pixel: "https://push.example.com/imp/dc3e7a05-e267-4a7a-88be-cec6e79f5d3d",
		Bid:   generator.PriceOrDefault(requestCtx.UserValue(string(Price)).(float64), generator.CPC),
	}

	aesearchResponse.Result.Listing = append(aesearchResponse.Result.Listing, aesearchResponseItem)

	return requestCtx.JSONResponse(&aesearchResponse, http.StatusOK)
}
