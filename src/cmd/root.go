package cmd

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"

	"github.com/KenshiTech/unchained/src/constants"
	"github.com/spf13/cobra"
)

var configPath string
var secretsPath string
var contextPath string
var printVersion bool

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "unchained",
	Short: "Unchained is the universal data validation and processing protocol",
	Long:  `Unchained is the universal data validation and processing protocol`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Println(constants.Version)
		} else {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.unchained.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
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
