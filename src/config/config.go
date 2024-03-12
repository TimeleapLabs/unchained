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
	Config.SetDefault("pos.address", "0xdA36f22C0Dab89CA0884489Bf3a00c0C27cEDbec")
	Config.SetDefault("pos.base", "1")
	Config.SetDefault("pos.nft", "40000000000000000000000")
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
