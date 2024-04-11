package cmd

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/ethereum"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/transport/server"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket"

	"github.com/spf13/cobra"
)

var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Start(config.App.System.Log)
		log.Logger.
			With("Mode", "Broker").
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained")

		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		err = ethereum.InitClientIdentity()
		if err != nil {
			panic(err)
		}

		bls.InitClientIdentity()

		ethRPC := ethereum.New()
		pos.New(ethRPC)

		server.New(
			websocket.WithWebsocket(),
		)
	},
}

func init() {
	rootCmd.AddCommand(brokerCmd)
}
