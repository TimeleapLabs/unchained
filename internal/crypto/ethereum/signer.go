package ethereum

import (
	"crypto/ecdsa"
)

type EvmSigner struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
	Address    string
}
