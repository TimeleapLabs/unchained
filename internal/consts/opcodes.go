package consts

type OpCode byte

// OpCodes of events.
const (
	OpCodeFeedback OpCode = iota
	OpCodeError    OpCode = 1

	OpCodeSubscribe   OpCode = 2
	OpCodeUnSubscribe OpCode = 3
	OpCodeMessage     OpCode = 4
	OpCodeBroadcast   OpCode = 5

	OpCodeRegisterWorker OpCode = 6
	OpCodeWorkerOverload OpCode = 7

	OpCodeRPCRequest  OpCode = 8
	OpCodeRPCResponse OpCode = 9
)
