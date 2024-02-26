/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/gql"
	"github.com/KenshiTech/unchained/net"
	"github.com/KenshiTech/unchained/plugins/logs"
	"github.com/KenshiTech/unchained/plugins/uniswap"

	"github.com/spf13/cobra"
)

// brokerCmd represents the broker command
var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Run the Unchained client in broker mode",
	Long:  `Run the Unchained client in broker mode`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig(configPath, secretsPath)
		db.Start()
		ethereum.Start()
		uniswap.Setup()
		logs.Setup()
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
