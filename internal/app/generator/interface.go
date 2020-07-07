package generator

type (
	IGenerator interface {
		// PriceOrDefault returns passed price if not zero or generate value for passed price model
		PriceOrDefault(float64, PriceModel) float64
		URLOrDefault(string) string
		AdMarkup() string

		OpenRTBURL(OpenRTBHandler) string
	}
)
