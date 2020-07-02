package generator

type (
	IGenerator interface {
		PriceOrDefault(float64, PriceModel) float64
		URL(URLType) string
		AdMarkup() string
	}
)
