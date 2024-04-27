package consts

type Channels string

const (
	ChannelPriceReport       Channels = "unchained:price_report"
	ChannelEventLog          Channels = "unchained:event_log"
	ChannelCorrectnessReport Channels = "unchained:correctness_report"
)
