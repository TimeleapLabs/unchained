package handler

import (
	"github.com/TimeleapLabs/unchained/internal/app"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/spf13/cobra"
)

// broker represents the broker command.
var broker = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,
	Run: func(_ *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		log.Start(config.App.System.Log)
		app.Broker()
	},
}

// WithBrokerCmd appends the broker command to the root command.
func WithBrokerCmd(cmd *cobra.Command) {
	cmd.AddCommand(broker)
}

func init() {
	broker.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
