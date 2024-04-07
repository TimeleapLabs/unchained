package handler

import (
	"github.com/KenshiTech/unchained/internal/app"
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/spf13/cobra"
)

// consumer represents the consumer command.
var consumer = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client in consumer mode",
	Long:  `Run the Unchained client in consumer mode`,

	PreRun: func(cmd *cobra.Command, args []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
		config.App.Network.Bind = cmd.Flags().Lookup("graphql").Value.String()
	},

	Run: func(cmd *cobra.Command, args []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		log.Start(config.App.System.Log)
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
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)

	consumer.Flags().StringP(
		"graphql",
		"g",
		"127.0.0.1:8080",
		"The graphql server path to bind",
	)
}
