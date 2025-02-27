package consts

import "errors"

// TODO: This is a mess
// Error list of the project.
var (
	ErrInvalidConfig           = errors.New("conf.invalid")
	ErrMissingHello            = errors.New("hello.missing")
	ErrInternalError           = errors.New("internal_error")
	ErrInvalidSignature        = errors.New("signature.invalid")
	ErrNotSupportedDataset     = errors.New("dataset not supported")
	ErrNotSupportedInstruction = errors.New("instruction not supported")
	ErrCantLoadSecret          = errors.New("can't load secrets")
	ErrCantLoadConfig          = errors.New("can't load config")
	ErrCantWriteSecret         = errors.New("can't write secrets")
	ErrTokenNotSupported       = errors.New("token not supported")
	ErrEventNotSupported       = errors.New("event not supported")
	ErrTopicNotSupported       = errors.New("topic not supported")
	ErrDataTooOld              = errors.New("data too old")
	ErrCantAggregateSignatures = errors.New("can't aggregate signatures")
	ErrCantRecoverSignature    = errors.New("can't recover signature")
	ErrClientNotFound          = errors.New("client not found")
	ErrSignatureNotfound       = errors.New("signature not found")
	ErrRecordNotfound          = errors.New("record not found")
	ErrCantLoadLastBlock       = errors.New("can't load last block")
	ErrDuplicateSignature      = errors.New("duplicate signature")
	ErrCrossPriceIsNotZero     = errors.New("cross price is not zero")
	ErrAlreadySynced           = errors.New("already synced")
	ErrCantSendRPCRequest      = errors.New("can't send rpc request")
	ErrCantReceiveRPCResponse  = errors.New("can't receive rpc response")
	ErrNoNewSigners            = errors.New("no new signers")
	ErrPluginNotFound          = errors.New("plugin not found")
	ErrFunctionNotFound        = errors.New("function not found")
	ErrOverloaded              = errors.New("worker overloaded")
	ErrNoWorker                = errors.New("no worker")
	ErrTimeout                 = errors.New("timeout")
	ErrInvalidPacket           = errors.New("invalid packet")
)
