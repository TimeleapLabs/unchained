package config_test

import (
	"os"
	"testing"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/stretchr/testify/assert"
)

var s = config.Secret{
	Address:    "n1",
	EvmAddress: "n2",
	SecretKey:  "n3",
	PublicKey:  "n4",
}

func TestSaveSecret(t *testing.T) {
	config.SecretFilePath = "./secret.yaml"
	err := s.Save()
	assert.Nil(t, err, "Should write successfully")

	err = os.Remove(config.SecretFilePath)
	assert.Nil(t, err, "Should delete successfully")
}
