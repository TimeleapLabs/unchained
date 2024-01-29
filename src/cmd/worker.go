/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/plugins/uniswap"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run the Unchained client in worker mode",
	Long:  `Run the Unchained client in worker mode`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig(configPath)
		uniswap.Work()
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

	workerCmd.Flags().StringVarP(&configPath, "config", "c", "./unchained.conf.yaml", "Config file")
	workerCmd.MarkFlagFilename("config", "yaml")
	workerCmd.MarkFlagRequired("config")
}
