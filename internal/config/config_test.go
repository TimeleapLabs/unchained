package config_test

import (
	"os"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/constants"
	"github.com/stretchr/testify/assert"
)

var s = config.Secret{
	Address:    "n1",
	EvmAddress: "n2",
	SecretKey:  "n3",
	PublicKey:  "n4",
}

func TestSaveSecret(t *testing.T) {
	err := s.Save()
	assert.Equal(t, constants.ErrCantWriteSecret, err, "Should return error because path of secret is not defined")

	config.SecretFilePath = "./secret.yaml"
	err = s.Save()
	assert.Nil(t, err, "Should write successfully")

	err = os.Remove(config.SecretFilePath)
	assert.Nil(t, err, "Should delete successfully")
}
