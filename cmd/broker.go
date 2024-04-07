package cmd

import (
	"github.com/KenshiTech/unchained/app"
	"github.com/spf13/cobra"
)

var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,
	Run: func(cmd *cobra.Command, args []string) {
		app.Broker()
	},
}

func init() {
	rootCmd.AddCommand(brokerCmd)
}
