package consts

// Channels represent all channels which can be subscribed by consumers.
type Channels string

const (
	ChannelPriceReport       Channels = "unchained:price_report"
	ChannelEventLog          Channels = "unchained:event_log"
	ChannelCorrectnessReport Channels = "unchained:correctness_report"
	ChannelFrostSigner       Channels = "unchained:frost:handshake"
)
