package handler

import (
	"github.com/TimeleapLabs/unchained/internal/app"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/spf13/cobra"
)

// worker represents the worker command.
var worker = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.Broker.URI = cmd.Flags().Lookup("broker").Value.String()
	},

	Run: func(cmd *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		utils.SetupLogger(config.App.System.Log)
		app.Worker(cmd.Context())
	},
}

// WithWorkerCmd appends the worker command to the root command.
func WithWorkerCmd(cmd *cobra.Command) {
	cmd.AddCommand(worker)
}

func init() {
	worker.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.timeleap.swiss",
		"Unchained broker to connect to",
	)
}
