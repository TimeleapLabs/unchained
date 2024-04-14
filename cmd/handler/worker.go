package handler

import (
	"github.com/KenshiTech/unchained/internal/app"
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/spf13/cobra"
)

// worker represents the worker command.
var worker = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
	},

	Run: func(_ *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		log.Start(config.App.System.Log)
		app.Worker()
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
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
