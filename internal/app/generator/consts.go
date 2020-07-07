package generator

type PriceModel string
type AdMarkupType string

type OpenRTBHandler string

const (
	CPM PriceModel = "CPM"
	CPC PriceModel = "CPC"
)

const (
	BURL OpenRTBHandler = "burl"
	NURL OpenRTBHandler = "nurl"
	LURL OpenRTBHandler = "lurl"
)

const (
	Banner AdMarkupType = "Banner"
	URL    AdMarkupType = "URL"
)
