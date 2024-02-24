package opcodes

// TODO: Should we have a Data opcode instead of PriceReport & EventLog?
const (
	Hello = iota
	KoskChallenge
	KoskResult

	RegisterConsumer

	Feedback
	Error

	PriceReport
	PriceReportBroadcast

	EventLog
	EventLogBroadcast
)
