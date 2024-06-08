package handler

import (
	"github.com/TimeleapLabs/unchained/internal/app"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/spf13/cobra"
)

// consumer represents the consumer command.
var consumer = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client in consumer mode",
	Long:  `Run the Unchained client in consumer mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
		config.App.Network.Bind = cmd.Flags().Lookup("graphql").Value.String()
	},
}

// postgresConsumer represents the consumer command which is run in Postgres Mode.
var postgresConsumer = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client for handling postgres in consumer mode",
	Long:  `Run the Unchained client for handling postgres in consumer mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
		config.App.Network.Bind = cmd.Flags().Lookup("graphql").Value.String()
	},

	Run: func(_ *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		utils.SetupLogger(config.App.System.Log)
		app.Consumer(app.PostgresConsumer)
	},
}

// schnorrConsumer represents the consumer command which is run in Schnorr Mode.
var schnorrConsumer = &cobra.Command{
	Use:   "schnorr",
	Short: "Run the Unchained client for handling schnorr in consumer mode",
	Long:  `Run the Unchained client for handling schnorr in consumer mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.BrokerURI = cmd.Flags().Lookup("broker").Value.String()
		config.App.Network.Bind = cmd.Flags().Lookup("graphql").Value.String()
	},

	Run: func(_ *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		utils.SetupLogger(config.App.System.Log)
		app.Consumer(app.SchnorrConsumer)
	},
}

// WithConsumerCmd append command of consumer to the root command.
func WithConsumerCmd(cmd *cobra.Command) {
	cmd.AddCommand(consumer)
	consumer.AddCommand(postgresConsumer)
	consumer.AddCommand(schnorrConsumer)
}

// init loads CLI flags of consumer command.
func init() {
	postgresConsumer.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)

	postgresConsumer.Flags().StringP(
		"graphql",
		"g",
		"127.0.0.1:8080",
		"The graphql server path to bind",
	)

	schnorrConsumer.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)

	schnorrConsumer.Flags().StringP(
		"graphql",
		"g",
		"127.0.0.1:8080",
		"The graphql server path to bind",
	)
}
