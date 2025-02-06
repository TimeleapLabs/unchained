package main

import (
	"fmt"
	"os"

	"github.com/TimeleapLabs/timeleap/cmd/handler"
	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/spf13/cobra"
)

// root represents the root command of application.
var root = &cobra.Command{
	Use:   "timeleap",
	Short: "Timeleap is the universal data validation and processing protocol",
	Long:  `Timeleap is the universal data validation and processing protocol`,
	Run: func(_ *cobra.Command, _ []string) {
		if config.App.System.PrintVersion {
			fmt.Println(consts.Version)
		} else {
			os.Exit(1)
		}
	},
}

func main() {
	handler.WithBrokerCmd(root)
	handler.WithConsumerCmd(root)
	handler.WithWorkerCmd(root)

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	root.Flags().BoolVarP(&config.App.System.PrintVersion, "version", "v", false, "Print the Timeleap version number and die")
	root.PersistentFlags().StringVarP(&config.App.System.ConfigPath, "config", "c", "./conf.yaml", "Config file")
	root.PersistentFlags().StringVarP(&config.App.System.SecretsPath, "secrets", "s", "./secrets.yaml", "Secrets file")
	root.PersistentFlags().BoolVarP(&config.App.System.AllowGenerateSecrets, "allow-generate-secrets", "a", false, "Allow to generate secrets file if not exists")
	root.PersistentFlags().StringVarP(&config.App.System.ContextPath, "context", "x", "./context", "Context DB")
	root.PersistentFlags().StringVarP(&config.App.System.Home, "home", "H", "./timeleap", "Timeleap Home")
}
