package opcodes

// TODO: We should have a Data opcode instead of PriceReport & EventLog
const (
	Hello = iota
	PriceReport
	Feedback
	KoskChallenge
	KoskResult
	Error
	RegisterConsumer
	ConsumeBroadcast
	EventLog
)
