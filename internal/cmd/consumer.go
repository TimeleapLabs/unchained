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
	"github.com/KenshiTech/unchained/pos"
	"github.com/KenshiTech/unchained/scheduler"
	"github.com/KenshiTech/unchained/scheduler/correctness"
	evmlogService "github.com/KenshiTech/unchained/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/service/uniswap"
	"github.com/KenshiTech/unchained/transport/client"
	"github.com/KenshiTech/unchained/transport/client/conn"
	"github.com/KenshiTech/unchained/transport/client/handler"
	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command.
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client in consumer mode",
	Long:  `Run the Unchained client in consumer mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Broker.URI = cmd.Flags().Lookup("broker").Value.String()
	},

	Run: func(cmd *cobra.Command, args []string) {
		log.Start(config.App.System.Log)
		log.Logger.
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained | Consumer")

		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		bls.InitClientIdentity()

		evmlogService := evmlogService.New()
		uniswapService := uniswapService.New()

		conn.Start()
		handler := handler.New(uniswapService, evmlogService)
		client.Consume(handler)

		ethereum.Start()
		pos.Start()
		db.Start()
		correctness.New()

		scheduler := scheduler.New(
			scheduler.WithEthLogs(),
			scheduler.WithUniswapEvents(),
		)

		scheduler.Start()
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)

	consumerCmd.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
