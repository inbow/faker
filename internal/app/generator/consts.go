package generator

type PriceModel string
type AdMarkupType string

type OpenRTBHandler string

const (
	Mile = 1000
)

const (
	CPM PriceModel = "CPM"
	CPC PriceModel = "CPC"
)

const (
	LossURL    OpenRTBHandler = "lurl"
	BiddingURL OpenRTBHandler = "burl"
	NoticeURL  OpenRTBHandler = "nurl"
)

const (
	Banner AdMarkupType = "Banner"
	URL    AdMarkupType = "URL"
)
