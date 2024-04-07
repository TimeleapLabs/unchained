package main

import (
	"fmt"
	"os"

	"github.com/KenshiTech/unchained/cmd/handler"
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/spf13/cobra"
)

// root represents the root command of application.
var root = &cobra.Command{
	Use:   "unchained",
	Short: "Unchained is the universal data validation and processing protocol",
	Long:  `Unchained is the universal data validation and processing protocol`,

	Run: func(cmd *cobra.Command, args []string) {
		if config.App.System.PrintVersion {
			fmt.Println(constants.Version)
		} else {
			os.Exit(1)
		}
	},
}

func main() {
	log.Start(config.App.System.Log)

	root.Flags().BoolVarP(&config.App.System.PrintVersion, "version", "v", false, "Print the Unchained version number and die")
	root.PersistentFlags().StringVarP(&config.App.System.ConfigPath, "config", "c", "./conf.yaml", "Config file")
	root.PersistentFlags().StringVarP(&config.App.System.SecretsPath, "secrets", "s", "./secrets.yaml", "Secrets file")
	root.PersistentFlags().StringVarP(&config.App.System.ContextPath, "context", "x", "./context", "Context DB")

	err := root.MarkFlagFilename("config", "yaml")
	if err != nil {
		log.Logger.Warn("no config flag")
	}

	err = root.MarkFlagRequired("config")
	if err != nil {
		log.Logger.Warn("no config flag")
	}

	handler.WithBrokerCmd(root)
	handler.WithConsumerCmd(root)
	handler.WithWorkerCmd(root)

	err = root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
