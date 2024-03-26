/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/app/consumer"
	"github.com/KenshiTech/unchained/config"

	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command.
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Run the Unchained client in consumer mode",
	Long:  `Run the Unchained client in consumer mode`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := config.Config.BindPFlag("broker.uri", cmd.Flags().Lookup("broker"))
		if err != nil {
			return err
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := consumer.NewConsumerApp(configPath, secretsPath)
		if err != nil {
			return err
		}
		return app.Run(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	consumerCmd.Flags().StringP(
		"broker",
		"b",
		"wss://shinobi.brokers.kenshi.io",
		"Unchained broker to connect to",
	)
}
