/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/db"
	"github.com/KenshiTech/unchained/src/ethereum"
	"github.com/KenshiTech/unchained/src/gql"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net"
	"github.com/KenshiTech/unchained/src/plugins/correctness"
	"github.com/KenshiTech/unchained/src/plugins/logs"
	"github.com/KenshiTech/unchained/src/plugins/uniswap"

	"github.com/spf13/cobra"
)

// brokerCmd represents the broker command.
var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Load(configPath, secretsPath)
		if err != nil {
			panic(err)
		}

		log.Start(config.App.System.Log)
		db.Start()
		correctness.New()
		ethereum.Start()
		uniswap.New()
		logs.New()
		gql.InstallHandlers()
		net.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(brokerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// brokerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// brokerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
