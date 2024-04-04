package cmd

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/persistence"
	"github.com/KenshiTech/unchained/pos"
	"github.com/KenshiTech/unchained/scheduler"
	correctnessService "github.com/KenshiTech/unchained/service/correctness"
	evmlogService "github.com/KenshiTech/unchained/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/service/uniswap"
	"github.com/KenshiTech/unchained/transport/client"
	"github.com/KenshiTech/unchained/transport/client/handler"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command.
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
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

		ethRPC := ethereum.New()
		pos := pos.New(ethRPC)
		badger := persistence.New(contextPath)

		correctnessService := correctnessService.New(ethRPC)
		evmLogService := evmlogService.New(ethRPC, pos)
		uniswapService := uniswapService.New(ethRPC, pos)

		scheduler.New(
			scheduler.WithEthLogs(evmLogService, ethRPC, badger),
			scheduler.WithUniswapEvents(uniswapService, ethRPC),
		)

		handler := handler.New(correctnessService, uniswapService, evmLogService)
		client.Consume(handler)
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
