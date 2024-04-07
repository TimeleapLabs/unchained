package handler

import (
	"github.com/KenshiTech/unchained/internal/app"
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/spf13/cobra"
)

// consumer represents the consumer command.
var consumer = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client in consumer mode",
	Long:  `Run the Unchained client in consumer mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
	},

	Run: func(cmd *cobra.Command, args []string) {
		app.Consumer()
	},
}

// WithConsumerCmd append command of consumer to the root command.
func WithConsumerCmd(cmd *cobra.Command) {
	cmd.AddCommand(consumer)

	consumer.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
