package generator

type (
	IGenerator interface {
		Price(PriceModel) float64
		URL(URLType) string
		AdMarkup() string
	}
)
