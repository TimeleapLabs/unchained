package consts

type Channels string

// Channels of broker to be watched for events.
const (
	ChannelPriceReport Channels = "timeleap:price_report"
	ChannelEventLog    Channels = "timeleap:event_log"
	ChannelAttestation Channels = "timeleap:attestation_report"
)
