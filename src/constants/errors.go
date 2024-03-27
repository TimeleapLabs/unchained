package constants

import "errors"

var (
	ErrCantSendPacket          = errors.New("socket.unreachable")
	ErrInvalidKosk             = errors.New("kosk.invalid")
	ErrInvalidConfig           = errors.New("conf.invalid")
	ErrKosk                    = errors.New("kosk.error")
	ErrMissingHello            = errors.New("hello.missing")
	ErrMissingKosk             = errors.New("kosk.missing")
	ErrInternalError           = errors.New("internal_error")
	ErrCantVerifyBls           = errors.New("cant_verify_bls")
	ErrInvalidSignature        = errors.New("signature.invalid")
	ErrNotSupportedDataset     = errors.New("dataset not supported")
	ErrNotSupportedInstruction = errors.New("instruction not supported")
)
