package main

import (
	"fmt"
	"os"

	"github.com/TimeleapLabs/unchained/cmd/handler"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/spf13/cobra"
)

// root represents the root command of application.
var root = &cobra.Command{
	Use:   "unchained",
	Short: "Unchained is the universal data validation and processing protocol",
	Long:  `Unchained is the universal data validation and processing protocol`,
	Run: func(_ *cobra.Command, _ []string) {
		if config.App.System.PrintVersion {
			fmt.Println(consts.Version)
		} else {
			os.Exit(1)
		}
	},
}

// Unchained entrypoint.
func main() {
	handler.WithBrokerCmd(root)
	handler.WithConsumerCmd(root)
	handler.WithWorkerCmd(root)

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init loads global CLI flags.
func init() {
	root.Flags().BoolVarP(&config.App.System.PrintVersion, "version", "v", false, "Print the Unchained version number and die")
	root.PersistentFlags().StringVarP(&config.App.System.ConfigPath, "config", "c", "./conf.yaml", "Config file")
	root.PersistentFlags().StringVarP(&config.App.System.SecretsPath, "secrets", "s", "./secrets.yaml", "Secrets file")
	root.PersistentFlags().BoolVarP(&config.App.System.AllowGenerateSecrets, "allow-generate-secrets", "a", false, "Allow to generate secrets file if not exists")
	root.PersistentFlags().StringVarP(&config.App.System.ContextPath, "context", "x", "./context", "Context DB")
}
