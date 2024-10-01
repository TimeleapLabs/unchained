package consts

// TODO: Should we have a Data opcode instead of PriceReport & EventLog?

type OpCode byte

const (
	OpCodeHello         OpCode = iota
	OpCodeKoskChallenge OpCode = 1
	OpCodeKoskResult    OpCode = 2

	OpCodeRegisterConsumer OpCode = 3

	OpCodeFeedback OpCode = 4
	OpCodeError    OpCode = 5

	OpCodeCorrectnessReport          OpCode = 10
	OpCodeCorrectnessReportBroadcast OpCode = 11

	OpCodeRegisterRPCFunction OpCode = 12

	OpCodeRPCRequest  OpCode = 13
	OpCodeRPCResponse OpCode = 14
)
