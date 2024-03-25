/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	clientIdentity "github.com/KenshiTech/unchained/crypto/client_identity"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/client"
	"github.com/KenshiTech/unchained/persistence"
	"github.com/KenshiTech/unchained/plugins/logs"
	"github.com/KenshiTech/unchained/plugins/uniswap"
	"github.com/KenshiTech/unchained/pos"

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
		var err error
		config.LoadConfig(configPath, secretsPath)
		log.Start()

		log.Logger.
			With("Version", constants.Version).
			With("Protocol", constants.ProtocolVersion).
			Info("Running Unchained")

		ethereum.Start()
		// initializes client identity
		{
			var ops []clientIdentity.Option
			// read configuration
			{
				// SecretKey
				if v := config.Secrets.GetString(constants.SecretKey); v != "" {
					ops = append(ops, clientIdentity.OptionWithSecretKey(v))
				}

				// Name
				if v := config.Config.GetString(constants.Name); v != "" {
					ops = append(ops, clientIdentity.OptionWithName(v))
				}

				// EVMWallet
				if v := config.Secrets.GetString(constants.EVMWallet); v != "" {
					ops = append(ops, clientIdentity.OptionWithEvmWallet(v))
				}
			}

			err = clientIdentity.Init(ops...)
			if err != nil {
				return err
			}

		}
		pos.Start()
		db.Start()
		client.StartClient()
		uniswap.Setup()
		uniswap.Start()
		logs.Setup()
		logs.Start()
		persistence.Start(contextPath)
		client.Listen()
		return nil
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
