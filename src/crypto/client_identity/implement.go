package clientidentity

import (
	"math/big"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/log"
	"github.com/btcsuite/btcutil/base58"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

// clientIdentity
type clientIdentity struct {
	// SecretKey or PrivateKey is used for signing messages.
	SecretKey      *big.Int
	PublicKey      *bls12381.G2Affine
	ShortPublicKey *bls12381.G1Affine

	Signer *bls.Signer

	// addr holds the address of the client
	addr address.Address

	config struct {

		// secretKey hex encoded of the private key.
		secretKey string
		evmWallet string
		name      string
	}
}

var _ ClientIdentity = &clientIdentity{}

// New creates a new client identity based on the provided configuration.
func New(ops ...Option) (ClientIdentity, error) {
	c := new(clientIdentity)
	// c.loadDefaults()
	for _, fn := range ops {
		err := fn(c)
		if err != nil {
			return nil, err
		}
	}
	err := c.init()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Init initializes the client identity as singleton pattern.
// Flat functions will be callable when the Init is called.
func Init(ops ...Option) error {
	ci, err := New(ops...)
	if err != nil {
		return err
	}
	Default = ci
	return nil
}

func (c *clientIdentity) init() error {
	var err error

	// calculate key
	if c.config.secretKey != "" {
		// read from config
		bs := base58.Decode(c.config.name)
		c.SecretKey = new(big.Int).SetBytes(bs)
		c.PublicKey = bls.GetPublicKey(c.SecretKey)
	} else {
		// generate a new pair of secret
		c.SecretKey, c.PublicKey, err = bls.GenerateKeyPair()
		if err != nil {
			return err
		}
	}

	c.ShortPublicKey = bls.GetShortPublicKey(c.SecretKey)

	pkBytes := c.PublicKey.Bytes()
	c.addr, err = address.NewAddress(pkBytes[:])
	if err != nil {
		return err
	}

	c.Signer = &bls.Signer{
		Name:           c.config.name,
		EvmWallet:      c.config.evmWallet,
		PublicKey:      c.PublicKey.Bytes(),
		ShortPublicKey: c.ShortPublicKey.Bytes(),
	}

	log.Logger.With(constants.Address, c.addr.String()).Info(constants.Unchained)

	return nil
}

func (c *clientIdentity) GetSecretKey() *big.Int {
	return c.SecretKey
}

func (c *clientIdentity) GetPublicKey() *bls12381.G2Affine {
	return c.PublicKey
}

func (c *clientIdentity) GetShortPublicKey() *bls12381.G1Affine {
	return c.ShortPublicKey
}

func (c *clientIdentity) GetSigner() *bls.Signer {
	return c.Signer
}
