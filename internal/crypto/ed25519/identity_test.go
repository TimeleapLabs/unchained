package ed25519

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/stretchr/testify/suite"
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
	signature, err := s.signer.Sign(TestMessage)
	s.NoError(err)

	isValid := s.signer.Verify(signature, TestMessage, s.signer.PublicKey)
	s.True(isValid)

	secondSignature, err := s.signer.Sign(SecondMessage)
	s.NoError(err)

	isValid = s.signer.Verify(signature, SecondMessage, s.signer.PublicKey)
	s.False(isValid)

	isValid = s.signer.Verify(secondSignature, TestMessage, s.signer.PublicKey)
	s.False(isValid)
}

func TestIdentitySuite(t *testing.T) {
	suite.Run(t, new(TestIdentitySuit))
}
