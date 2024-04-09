package cmd

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/db"
	"github.com/KenshiTech/unchained/internal/ethereum"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/pos"
	correctnessService "github.com/KenshiTech/unchained/internal/service/correctness"
	evmlogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
	"github.com/KenshiTech/unchained/internal/transport/server"
	"github.com/KenshiTech/unchained/internal/transport/server/gql"
	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client in consumer mode",
	Long:  `Run the Unchained client in consumer mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
		config.App.Network.Bind = cmd.Flags().Lookup("graphql").Value.String()
	},

	Run: func(cmd *cobra.Command, args []string) {
		log.Start(config.App.System.Log)
		log.Logger.
			With("Mode", "Consumer").
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained")

		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		bls.InitClientIdentity()

		ethRPC := ethereum.New()
		pos := pos.New(ethRPC)
		db.Start()

		correctnessService := correctnessService.New(ethRPC)
		evmLogService := evmlogService.New(ethRPC, pos)
		uniswapService := uniswapService.New(ethRPC, pos)

		conn.Start()

		handler := handler.NewConsumerHandler(correctnessService, uniswapService, evmLogService)
		client.NewRPC(handler)

		server.New(
			gql.WithGraphQL(),
		)
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

	consumerCmd.Flags().StringP(
		"graphql",
		"g",
		"127.0.0.1:8080",
		"The graphql server path to bind",
	)
}
