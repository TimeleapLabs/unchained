/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/constants"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/db"
	"github.com/KenshiTech/unchained/src/ethereum"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/client"
	"github.com/KenshiTech/unchained/src/persistence"
	"github.com/KenshiTech/unchained/src/plugins/logs"
	"github.com/KenshiTech/unchained/src/plugins/uniswap"
	"github.com/KenshiTech/unchained/src/pos"

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
		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		log.Start(config.App.System.Log)
		log.Logger.
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained")

		ethereum.Start()
		bls.InitClientIdentity()
		pos.Start()
		db.Start()
		client.StartClient()

		uniswap.Listen()
		logs.New()
		logs.Listen()

		persistence.Start(contextPath)
		client.Listen()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	workerCmd.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
