package opcodes

// TODO: Should we have a Data opcode instead of PriceReport & EventLog?

type OpCode byte

const (
	Hello         OpCode = iota
	KoskChallenge OpCode = 1
	KoskResult    OpCode = 2

	RegisterConsumer OpCode = 3

	Feedback OpCode = 4
	Error    OpCode = 5

	PriceReport          OpCode = 6
	PriceReportBroadcast OpCode = 7

	EventLog          OpCode = 8
	EventLogBroadcast OpCode = 9

	CorrectnessReport          OpCode = 10
	CorrectnessReportBroadcast OpCode = 11
)
