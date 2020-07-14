package generator

type PriceModel string
type OpenRTBHandler string

const (
	Mile = 1000
)

const (
	CPV PriceModel = "CPV"
	CPM PriceModel = "CPM"
	CPC PriceModel = "CPC"
)

const (
	LossURL    OpenRTBHandler = "loss_url"
	NoticeURL  OpenRTBHandler = "notice_url"
	BiddingURL OpenRTBHandler = "bidding_url"
)
