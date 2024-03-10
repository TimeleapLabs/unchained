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
	Config.SetDefault("broker.bind", "0.0.0.0:9123")
	Config.SetDefault("pos.chain", "arbitrum_sepolia")
	Config.SetDefault("pos.address", "0x08cE842914b7313E16DBA708ccF0418bE3dE05c6")
	Config.SetDefault("pos.base", "1")
	Config.SetDefault("pos.nft", "25000000000000000000000")
}

var Config *viper.Viper
var Secrets *viper.Viper

func init() {
	Config = viper.New()
	Secrets = viper.New()
}

func LoadConfig(ConfigFileName string, SecretsFileName string) {
	defaults()

	Config.SetConfigFile(ConfigFileName)
	err := Config.ReadInConfig()

	if err != nil {
		if os.IsNotExist(err) {
			Config.WriteConfig()
		} else {
			panic(err)
		}
	}

	Secrets.SetConfigFile(SecretsFileName)
	err = Secrets.MergeInConfig()

	if err != nil && os.IsExist(err) {
		panic(err)
	}

}
