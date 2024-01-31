package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.11.0-alpha.1"

var configPath string
var printVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unchained",
	Short: "Unchained is the universal data validation and processing protocol",
	Long:  `Unchained is the universal data validation and processing protocol`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Println(version)
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "Print the Unchained version number and die")

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./unchained.conf.yaml", "Config file")
	rootCmd.MarkFlagFilename("config", "yaml")
	rootCmd.MarkFlagRequired("config")
}
