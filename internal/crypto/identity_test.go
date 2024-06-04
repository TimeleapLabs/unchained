package crypto

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/assert"
)

const SamplePrivateKey = "3b885a8a8f043724abfa865eccd38f536887d9ea1c08a742720e810f38a86872"

func TestEvmSignerWithoutGeneratePrivateKey(t *testing.T) {
	utils.SetupLogger("info")
	config.App.Secret.EvmPrivateKey = SamplePrivateKey

	InitMachineIdentity(
		WithEvmSigner(),
	)

	assert.Equal(t, config.App.Secret.EvmPrivateKey, SamplePrivateKey)
}

func TestEvmSignerWithGeneratePrivateKey(t *testing.T) {
	utils.SetupLogger("info")
	config.App.Secret.EvmPrivateKey = ""

	InitMachineIdentity(
		WithEvmSigner(),
	)

	assert.NotEmpty(t, config.App.Secret.EvmPrivateKey)
}
