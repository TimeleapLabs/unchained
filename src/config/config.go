package config

import (
	"os"
	"path"
	"runtime"

	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/ilyakaznacheev/cleanenv"
)

var App Config
var SecretFilePath string

func Load(configPath, secretPath string) error {
	if configPath == "" {
		_, b, _, _ := runtime.Caller(0)
		configPath = path.Join(b, "../..", "./config.yaml")
	}

	if secretPath != "" {
		SecretFilePath = secretPath
		err := cleanenv.ReadConfig(secretPath, App.Secret)
		if err != nil {
			log.Err(err).Msg("Can't read secret file")
			return constants.ErrCantLoadSecret
		}
	}

	err := cleanenv.ReadConfig(configPath, App)
	if err != nil {
		log.Err(err).Msg("Can't read config file")
		return constants.ErrCantLoadConfig
	}

	err = cleanenv.ReadEnv(App)
	if err != nil {
		return err
	}

	return nil
}

func (s *Secret) Save() error {
	yamlData, err := yaml.Marshal(&s)
	if err != nil {
		log.Err(err).Msg("Can't marshal to yaml")
		return constants.ErrCantWriteSecret
	}

	if SecretFilePath == "" {
		log.Err(err).Msg("SecretFilePath is not defined")
		return constants.ErrCantWriteSecret
	}

	err = os.WriteFile(SecretFilePath, yamlData, 0600)
	if err != nil {
		log.Err(err).Msg("Can't write secret file")
		return constants.ErrCantWriteSecret
	}

	return nil
}
