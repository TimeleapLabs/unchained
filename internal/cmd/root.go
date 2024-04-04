package cmd

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"

	"github.com/KenshiTech/unchained/constants"
	"github.com/spf13/cobra"
)

var configPath string
var secretsPath string
var contextPath string
var printVersion bool

var rootCmd = &cobra.Command{
	Use:   "unchained",
	Short: "Unchained is the universal data validation and processing protocol",
	Long:  `Unchained is the universal data validation and processing protocol`,

	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Println(constants.Version)
		} else {
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "Print the Unchained version number and die")

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./conf.yaml", "Config file")
	rootCmd.PersistentFlags().StringVarP(&secretsPath, "secrets", "s", "./secrets.yaml", "Secrets file")
	rootCmd.PersistentFlags().StringVarP(&contextPath, "context", "x", "./context", "Context DB")
	err := rootCmd.MarkFlagFilename("config", "yaml")
	if err != nil {
		log.Warn("no config flag")
	}

	err = rootCmd.MarkFlagRequired("config")
	if err != nil {
		log.Warn("no config flag")
	}
}
