/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/persistence"
	"github.com/KenshiTech/unchained/pos"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command.
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Broker.URI = cmd.Flags().Lookup("broker").Value.String()
	},

	Run: func(cmd *cobra.Command, args []string) {
		log.Start(config.App.System.Log)
		log.Logger.
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained | Worker")

		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		bls.InitClientIdentity()

		ethereum.Start()
		pos.Start()
		db.Start()

		persistence.Start(contextPath)
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	workerCmd.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
