package generator

type (
	IGenerator interface {
		PriceOrDefault(float64, PriceModel) float64
		URLOrDefault(string) string
		AdMarkup() string

		OpenRTBURL(OpenRTBHandler) string
	}
)
