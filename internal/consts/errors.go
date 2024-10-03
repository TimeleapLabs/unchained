package consts

import "errors"

var (
	ErrInvalidKosk             = errors.New("kosk.invalid")
	ErrInvalidConfig           = errors.New("conf.invalid")
	ErrKosk                    = errors.New("kosk.error")
	ErrMissingHello            = errors.New("hello.missing")
	ErrMissingKosk             = errors.New("kosk.missing")
	ErrInternalError           = errors.New("internal_error")
	ErrCantVerifyBls           = errors.New("cant_verify_bls")
	ErrInvalidSignature        = errors.New("signature.invalid")
	ErrNotSupportedInstruction = errors.New("instruction not supported")
	ErrCantLoadSecret          = errors.New("can't load secrets")
	ErrCantLoadConfig          = errors.New("can't load config")
	ErrCantWriteSecret         = errors.New("can't write secrets")
	ErrTopicNotSupported       = errors.New("topic not supported")
	ErrClientNotFound          = errors.New("client not found")
	ErrDuplicateSignature      = errors.New("duplicate signature")
	ErrCantSendRPCRequest      = errors.New("can't send rpc request")
	ErrCantReceiveRPCResponse  = errors.New("can't receive rpc response")
)
