package config

import "github.com/spf13/viper"

func defaults() {
	viper.SetDefault("rpc.ethereum", "https://eth.llamarpc.com")
}

func LoadConfig(FileName string) {
	defaults()

	viper.SetConfigFile(FileName)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}
