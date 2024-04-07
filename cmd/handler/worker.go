package handler

import (
	"github.com/KenshiTech/unchained/internal/app"
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/spf13/cobra"
)

// worker represents the worker command.
var worker = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
	},

	Run: func(cmd *cobra.Command, args []string) {
		app.Worker()
	},
}

// WithWorkerCmd appends the worker command to the root command.
func WithWorkerCmd(cmd *cobra.Command) {
	cmd.AddCommand(worker)

	worker.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
