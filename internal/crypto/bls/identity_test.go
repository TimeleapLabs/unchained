package bls

import (
	"crypto/rand"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	bls12381_fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	testData       = []byte("HELLO hello")
	secondTestData = []byte("HELLO hello 2")
)

type BlsIdentityTestSuite struct {
	suite.Suite
	identity *Signer
}

func (s *BlsIdentityTestSuite) SetupTest() {
	utils.SetupLogger("info")

	config.App.System.AllowGenerateSecrets = true
	s.identity = NewIdentity()
}

func (s *BlsIdentityTestSuite) TestSign() {
	signedData, err := s.identity.Sign(testData)
	assert.NoError(s.T(), err)

	publicKey := s.identity.PublicKey.Bytes()

	s.Run("Should verify the message correctly", func() {
		isVerified, err := s.identity.Verify(signedData, testData, publicKey[:])
		assert.NoError(s.T(), err)
		assert.True(s.T(), isVerified)
	})

	s.Run("Should not verify the message because of wrong data case", func() {
		isVerified, err := s.identity.Verify(signedData, secondTestData, publicKey[:])
		assert.NoError(s.T(), err)
		assert.False(s.T(), isVerified)
	})

	s.Run("Should not verify the message because of wrong data case", func() {
		signedSecondData, err := s.identity.Sign(secondTestData)
		assert.NoError(s.T(), err)

		isVerified, err := s.identity.Verify(signedSecondData, testData, publicKey[:])
		assert.NoError(s.T(), err)
		assert.False(s.T(), isVerified)
	})

	s.Run("Should not verify the message because publicKey is not belong to this message", func() {
		_, _, _, g2Aff := bls12381.Generators()
		g2Order := bls12381_fr.Modulus()
		sk, err := rand.Int(rand.Reader, g2Order)
		if err != nil {
			panic(err)
		}

		pk := new(bls12381.G2Affine).ScalarMultiplication(&g2Aff, sk).Bytes()

		isVerified, err := s.identity.Verify(signedData, secondTestData, pk[:])
		assert.NoError(s.T(), err)
		assert.False(s.T(), isVerified)
	})
}

func TestBlsIdentitySuite(t *testing.T) {
	suite.Run(t, new(BlsIdentityTestSuite))
}
