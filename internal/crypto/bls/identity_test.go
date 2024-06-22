package bls

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/stretchr/testify/suite"
	"testing"
)

const SamplePrivateKey = "3b885a8a8f043724abfa865eccd38f536887d9ea1c08a742720e810f38a86872"

var (
	TestMessage   = []byte("test message")
	SecondMessage = []byte("second message")
)

type TestIdentitySuit struct {
	suite.Suite
	signer *Signer
}

func (s *TestIdentitySuit) SetupTest() {
	config.App.System.AllowGenerateSecrets = true
	config.App.Secret.EvmPrivateKey = SamplePrivateKey

	s.signer = NewIdentity()
}

func (s *TestIdentitySuit) TestVerifyIdentity() {
	signature, messageHash := s.signer.Sign(TestMessage)
	isValid, err := s.signer.Verify(signature, messageHash, *s.signer.PublicKey)
	s.NoError(err)
	s.True(isValid)

	secondSignature, secondMessageHash := s.signer.Sign(SecondMessage)

	isValid, err = s.signer.Verify(signature, secondMessageHash, *s.signer.PublicKey)
	s.NoError(err)
	s.False(isValid)

	isValid, err = s.signer.Verify(secondSignature, messageHash, *s.signer.PublicKey)
	s.NoError(err)
	s.False(isValid)
}

func TestIdentitySuite(t *testing.T) {
	suite.Run(t, new(TestIdentitySuit))
}
