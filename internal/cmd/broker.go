/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/transport/server"
	"github.com/KenshiTech/unchained/transport/server/gql"
	"github.com/KenshiTech/unchained/transport/server/websocket"

	"github.com/spf13/cobra"
)

var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Start(config.App.System.Log)
		log.Logger.
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained | Broker")

		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		db.Start()
		ethereum.Start()

		server.New(
			gql.WithGraphQL(),
			websocket.WithWebsocket(),
		)
	},
}

func init() {
	rootCmd.AddCommand(brokerCmd)
}
