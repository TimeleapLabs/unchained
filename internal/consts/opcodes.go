package consts

type OpCode byte

// OpCodes of events.
const (
	OpCodeHello         OpCode = iota
	OpCodeKoskChallenge OpCode = 1
	OpCodeKoskResult    OpCode = 2

	OpCodeRegisterConsumer OpCode = 3

	OpCodeFeedback OpCode = 4
	OpCodeError    OpCode = 5

	OpCodeAttestation          OpCode = 6
	OpCodeAttestationBroadcast OpCode = 7

	OpCodeRegisterWorker OpCode = 8

	OpCodeRPCRequest  OpCode = 9
	OpCodeRPCResponse OpCode = 10
)
