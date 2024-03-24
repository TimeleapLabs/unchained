package config

import (
	"os"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/spf13/viper"
)

func defaults() {
	Config.SetDefault("name", petname.Generate(3, "-"))
	Config.SetDefault("log", "info")
	Config.SetDefault("rpc.ethereum", "https://ethereum.publicnode.com")
	Config.SetDefault("rpc.arbitrum_sepolia", "https://sepolia-rollup.arbitrum.io/rpc")
	Config.SetDefault("broker.bind", "0.0.0.0:9123")
	Config.SetDefault("pos.chain", "arbitrum_sepolia")
	Config.SetDefault("pos.address", "0x965e364987356785b7E89e2Fe7B70f5E5107332d")
	Config.SetDefault("pos.base", "1")
}

var Config *viper.Viper
var Secrets *viper.Viper

func init() {
	Config = viper.New()
	Secrets = viper.New()
}

func LoadConfig(configFileName string, secretsFileName string) {
	defaults()

	Config.SetConfigFile(configFileName)
	err := Config.ReadInConfig()

	if err != nil {
		isNotExist := os.IsNotExist(err)
		if !isNotExist {
			panic(err)
		}

		err = Config.WriteConfig()
		if err != nil {
			panic(err)
		}
	}

	Secrets.SetConfigFile(secretsFileName)
	err = Secrets.MergeInConfig()

	if err != nil && os.IsExist(err) {
		panic(err)
	}
}
