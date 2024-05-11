package config

import (
	"os"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"

	pureLog "log"

	"gopkg.in/yaml.v3"

	"github.com/ilyakaznacheev/cleanenv"
)

// App holds global configs of application.
var App Config

// SecretFilePath holds the path of secret file.
var SecretFilePath string

// Load function read config and secret file and override them if env variable provided.
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
		pureLog.Println("Can't load secrets")
		// return constants.ErrCantLoadSecret
	}

	err = cleanenv.ReadConfig(configPath, &App)
	if err != nil {
		return consts.ErrCantLoadConfig
	}

	err = cleanenv.ReadEnv(&App)
	if err != nil {
		return err
	}

	return nil
}

// Save function write app's secret values to secret file.
func (s *Secret) Save() error {
	yamlData, err := yaml.Marshal(&s)
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't marshal secrets to yaml")
		return consts.ErrCantWriteSecret
	}

	if SecretFilePath == "" {
		SecretFilePath = "./secrets.yaml" // #nosec G101
	}

	err = os.WriteFile(SecretFilePath, yamlData, 0600)
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't write secret file")
		return consts.ErrCantWriteSecret
	}

	return nil
}
