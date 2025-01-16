package consts

type OpCode byte

// OpCodes of events.
const (
	OpCodeHello OpCode = iota

	OpCodeRegisterConsumer OpCode = 1

	OpCodeFeedback OpCode = 2
	OpCodeError    OpCode = 3

	OpCodeAttestation          OpCode = 4
	OpCodeAttestationBroadcast OpCode = 5

	OpCodeRegisterWorker OpCode = 6
	OpCodeWorkerOverload OpCode = 7

	OpCodeRPCRequest  OpCode = 8
	OpCodeRPCResponse OpCode = 9
)
