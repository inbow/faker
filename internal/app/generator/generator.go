package generator

import (
	"net/url"
	"strings"

	rd "github.com/Pallinder/go-randomdata"
)

type (
	Generator struct{}
)

func New() IGenerator {
	return &Generator{}
}

func (g *Generator) Price(priceModel PriceModel) float64 {
	price := rd.Decimal(5, 9)
	price /= 10 // Can't generate from 0.5 to 0.9 with this library :(

	if priceModel == CPC {
		price /= 1000
	}

	return price
}

func (g *Generator) URL(urlType URLType) string {
	resultedURL := &url.URL{}
	resultedURL.Scheme = "https"
	resultedURL.Host = strings.ToLower(rd.Letters(5)) + "." + strings.ToLower(rd.Letters(3))
	resultedURL.Path = string(urlType)

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

func (g *Generator) AdMarkup() string {
	return `<html><head></head><body>Hello World</body></html>`
}
