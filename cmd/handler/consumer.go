package handler

import (
	"github.com/TimeleapLabs/timeleap/internal/app"
	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/spf13/cobra"
)

// consumer represents the consumer command.
var consumer = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Timeleap client in consumer mode",
	Long:  `Run the Timeleap client in consumer mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.Broker.URI = cmd.Flags().Lookup("broker").Value.String()
		config.App.Network.Bind = cmd.Flags().Lookup("graphql").Value.String()
	},

	Run: func(_ *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		utils.SetupLogger(config.App.System.Log)
		app.Consumer()
	},
}

// WithConsumerCmd append command of consumer to the root command.
func WithConsumerCmd(cmd *cobra.Command) {
	cmd.AddCommand(consumer)
}

func init() {
	consumer.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.timeleap.swiss",
		"Timeleap broker to connect to",
	)

	consumer.Flags().StringP(
		"graphql",
		"g",
		"127.0.0.1:8080",
		"The graphql server path to bind",
	)
}
