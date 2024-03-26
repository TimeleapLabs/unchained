package worker

import (
	"context"

	"github.com/KenshiTech/unchained/app"
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
)

// WorkerApp implements the App to be used in CMD layer.
type WorkerApp struct {
	config struct {
		configPath  string
		secretsPath string
		contextPath string
	}
}

func NewWorkerApp(configPath, secretsPath, contextPath string) (app.App, error) {
	a := new(WorkerApp)
	a.config.configPath = configPath
	a.config.secretsPath = secretsPath
	a.config.contextPath = contextPath
	return a, nil
}

func (app *WorkerApp) Run(ctx context.Context) error {
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

	//Initialize ProofOfStake
	err = pos.Start()
	if err != nil {
		return err
	}

	//? WIP
	db.Start()
	client.StartClient()
	uniswap.Setup()
	uniswap.Start()
	logs.Setup()
	logs.Start()
	persistence.Start(app.config.contextPath)
	client.Listen()
	return nil
}
