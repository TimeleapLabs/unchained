package consumer

import (
	"context"

	"github.com/KenshiTech/unchained/app"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/consumers"
	clientIdentity "github.com/KenshiTech/unchained/crypto/client_identity"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/client"
	"github.com/KenshiTech/unchained/plugins/correctness"
	"github.com/KenshiTech/unchained/plugins/logs"
	"github.com/KenshiTech/unchained/plugins/uniswap"
	"github.com/KenshiTech/unchained/pos"
)

// ConsumerApp implements the App to be used in CMD layer.
type ConsumerApp struct {
	config struct {
		configPath  string
		secretsPath string
	}
}

func NewConsumerApp(configPath, secretsPath string) (app.App, error) {
	a := new(ConsumerApp)
	a.config.configPath = configPath
	a.config.secretsPath = secretsPath
	return a, nil
}

func (app *ConsumerApp) Run(ctx context.Context) error {
	var err error

	// load configuration
	err = config.LoadConfig(app.config.configPath, app.config.secretsPath)
	if err != nil {
		return err
	}

	// Init Logger
	err = log.Start()
	if err != nil {
		return err
	}
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

	// Initialize ProofOfStake module
	err = pos.Start()
	if err != nil {
		return err
	}

	db.Start()
	uniswap.Setup()
	correctness.Setup()
	logs.Setup()
	client.StartClient()
	consumers.StartConsumer()
	client.Listen()
	return nil
}
