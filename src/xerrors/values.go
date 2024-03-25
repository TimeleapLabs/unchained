package xerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPacket           = errors.New("packet.invalid")
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

	ErrNilArgs          = func(v interface{}) error { return fmt.Errorf("nil args: %v", v) }
	ErrConnectionFailed = func(v interface{}) error { return fmt.Errorf("connection failed: %v", v) }
)
