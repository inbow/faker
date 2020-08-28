package generator

import (
	"net/url"
	"strings"

	rd "github.com/Pallinder/go-randomdata"

	"github.com/oxyd-io/faker/internal/config"
)

func PriceOrDefault(price float64, priceModel PriceModel) float64 {
	if price != 0 {
		return price
	}

	price = rd.Decimal(5, 9) / 10 // Can't generate from 0.5 to 0.9 with this library :(

	if priceModel == CPC || priceModel == CPV {
		price /= Mile
	}

	return price
}

func OpenRTBURL(config *config.AppConfig, handler OpenRTBHandler) string {
	resultedURL := &url.URL{}
	resultedURL.Scheme = "http"
	resultedURL.Host = config.HTTP.Domain
	resultedURL.Path = "api/v1/openrtb/" + string(handler)

	query := resultedURL.Query()
	query.Add("ai", "${AUCTION_ID}")
	query.Add("ap", "${AUCTION_PRICE}")
	query.Add("ac", "${AUCTION_CURRENCY}")
	query.Add("abi", "${AUCTION_BID_ID}")
	query.Add("aii", "${AUCTION_IMP_ID}")
	query.Add("asi", "${AUCTION_SEAT_ID}")
	query.Add("aai", "${AUCTION_AD_ID}")

	unescapedQuery, _ := url.QueryUnescape(query.Encode())
	resultedURL.RawQuery = unescapedQuery

	return resultedURL.String()
}

func URLOrDefault(url string) string {
	if url != "" {
		return url
	}

	return "http://" + strings.ToLower(rd.Letters(5)) + "." + strings.ToLower(rd.Letters(3))
}
