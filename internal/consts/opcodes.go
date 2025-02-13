package consts

type OpCode byte

// OpCodes of events.
const (
	OpCodeFeedback OpCode = iota
	OpCodeError    OpCode = 1

	OpCodeSubscribe OpCode = 2
	OpCodeMessage   OpCode = 3
	OpCodeBroadcast OpCode = 4

	OpCodeRegisterWorker OpCode = 5
	OpCodeWorkerOverload OpCode = 6

	OpCodeRPCRequest  OpCode = 7
	OpCodeRPCResponse OpCode = 8
)
