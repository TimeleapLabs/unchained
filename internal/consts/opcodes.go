package consts

type OpCode byte

// OpCodes of events.
const (
	OpCodeFeedback OpCode = iota
	OpCodeError    OpCode = 1

	OpCodeRegisterConsumer OpCode = 2

	OpCodeMessage OpCode = 3

	OpCodeRegisterWorker OpCode = 4
	OpCodeWorkerOverload OpCode = 5

	OpCodeRPCRequest  OpCode = 6
	OpCodeRPCResponse OpCode = 7
)
