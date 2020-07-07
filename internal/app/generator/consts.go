package generator

type PriceModel string
type URLType string

const (
	Mile = 1000
)

const (
	CPM PriceModel = "CPM"
	CPC PriceModel = "CPC"
)

const (
	LossURL         URLType = "openrtb/lurl"
	BiddingURL      URLType = "openrtb/burl"
	NotificationURL URLType = "openrtb/nurl"
)
