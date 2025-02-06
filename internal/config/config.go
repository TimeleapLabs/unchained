package config

import (
	"fmt"
	"os"

	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/utils"

	pureLog "log"

	"gopkg.in/yaml.v3"

	"github.com/ilyakaznacheev/cleanenv"
)

// App is holding all configuration.
var App Config
var SecretFilePath string

// Load loads configuration from files.
func Load(configPath, secretPath string) error {
	if configPath == "" {
		configPath = "./config.yaml"
	}

	if secretPath == "" {
		secretPath = "./secrets.yaml" // #nosec G101
	}

	SecretFilePath = secretPath
	err := cleanenv.ReadConfig(secretPath, &App.Secret)
	if err != nil {
		pureLog.Println("Cannot load secrets")
		// return constants.ErrCantLoadSecret
	}

	err = cleanenv.ReadConfig(configPath, &App)
	if err != nil {
		return fmt.Errorf("%w: %w", consts.ErrCantLoadConfig, err)
	}

	err = cleanenv.ReadEnv(&App)
	if err != nil {
		return err
	}

	return nil
}

// Save saves secret configurations to file.
func (s *Secret) Save() error {
	yamlData, err := yaml.Marshal(&s)
	if err != nil {
		utils.Logger.With("Error", err).Error("Cannot marshal secrets to yaml")
		return consts.ErrCantWriteSecret
	}

	if SecretFilePath == "" {
		SecretFilePath = "./secrets.yaml" // #nosec G101
	}

	err = os.WriteFile(SecretFilePath, yamlData, 0600)
	if err != nil {
		utils.Logger.With("Error", err).Error("Cannot write secret file")
		return consts.ErrCantWriteSecret
	}

	return nil
}
