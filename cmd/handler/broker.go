package handler

import (
	"github.com/TimeleapLabs/unchained/internal/app"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/spf13/cobra"
)

// broker represents the broker command.
var broker = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,

	PreRun: func(cmd *cobra.Command, _ []string) {
		config.App.Network.CertFile = cmd.Flags().Lookup("cert-file").Value.String()
		config.App.Network.KeyFile = cmd.Flags().Lookup("key-file").Value.String()
	},

	Run: func(_ *cobra.Command, _ []string) {
		err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
		if err != nil {
			panic(err)
		}

		utils.SetupLogger(config.App.System.Log)
		app.Broker()
	},
}

// WithBrokerCmd appends the broker command to the root command.
func WithBrokerCmd(cmd *cobra.Command) {
	cmd.AddCommand(broker)
}

func init() {
	broker.Flags().StringP(
		"cert-file",
		"C",
		"",
		"TLS certificate file",
	)

	broker.Flags().StringP(
		"key-file",
		"k",
		"",
		"TLS key file",
	)
}
