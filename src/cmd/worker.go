/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/app/worker"
	"github.com/KenshiTech/unchained/config"

	"github.com/spf13/cobra"
)

// workerCmd represents the worker command.
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := config.Config.BindPFlag("broker.uri", cmd.Flags().Lookup("broker"))
		if err != nil {
			return err
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := worker.NewWorkerApp(configPath, secretsPath, contextPath)
		if err != nil {
			return err
		}
		return app.Run(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	workerCmd.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
