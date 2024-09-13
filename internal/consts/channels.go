package consts

type Channels string

// Channels of broker to be watched for events
const (
	ChannelPriceReport       Channels = "unchained:price_report"
	ChannelEventLog          Channels = "unchained:event_log"
	ChannelCorrectnessReport Channels = "unchained:correctness_report"
)
