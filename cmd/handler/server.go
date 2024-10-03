package handler

import (
	"github.com/TimeleapLabs/unchained/internal/app"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/spf13/cobra"
)

// server represents the server command.
var server = &cobra.Command{
	Use:   "server",
	Short: "Run the Unchained app",
	Long:  `Run the Unchained app`,

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

// WithServerCmd appends the server command to the root command.
func WithServerCmd(cmd *cobra.Command) {
	cmd.AddCommand(server)
}

func init() {
	server.Flags().StringP(
		"cert-file",
		"C",
		"",
		"TLS certificate file",
	)

	server.Flags().StringP(
		"key-file",
		"k",
		"",
		"TLS key file",
	)
}
