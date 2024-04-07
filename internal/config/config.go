package config

import (
	"os"
	"path"
	"runtime"

	pureLog "log"

	"github.com/KenshiTech/unchained/internal/log"

	"github.com/KenshiTech/unchained/internal/constants"
	"gopkg.in/yaml.v3"

	"github.com/ilyakaznacheev/cleanenv"
)

var App Config
var SecretFilePath string

func Load(configPath, secretPath string) error {
	_, b, _, _ := runtime.Caller(0)

	if configPath == "" {
		configPath = path.Join(b, "../..", "./config.yaml")
	}

	if secretPath == "" {
		secretPath = path.Join(b, "../..", "./secrets.yaml")
	}

	SecretFilePath = secretPath
	err := cleanenv.ReadConfig(secretPath, &App.Secret)
	if err != nil {
		pureLog.Println("Can't load secrets")
		// return constants.ErrCantLoadSecret
	}

	err = cleanenv.ReadConfig(configPath, &App)
	if err != nil {
		return constants.ErrCantLoadConfig
	}

	err = cleanenv.ReadEnv(&App)
	if err != nil {
		return err
	}

	return nil
}

func (s *Secret) Save() error {
	yamlData, err := yaml.Marshal(&s)
	if err != nil {
		log.Logger.With("Error", err).Error("Can't marshal secrets to yaml")
		return constants.ErrCantWriteSecret
	}

	if SecretFilePath == "" {
		log.Logger.With("Error", err).Error("SecretFilePath is not defined")
		return constants.ErrCantWriteSecret
	}

	err = os.WriteFile(SecretFilePath, yamlData, 0600)
	if err != nil {
		log.Logger.With("Error", err).Error("Can't write secret file")
		return constants.ErrCantWriteSecret
	}

	return nil
}
