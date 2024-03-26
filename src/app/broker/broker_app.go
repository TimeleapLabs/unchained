package broker

import (
	"context"

	"github.com/KenshiTech/unchained/app"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/gql"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net"
	"github.com/KenshiTech/unchained/plugins/correctness"
	"github.com/KenshiTech/unchained/plugins/logs"
	"github.com/KenshiTech/unchained/plugins/uniswap"
)

// BrokerApp implements the App to be used in CMD layer.
type BrokerApp struct {
	config struct {
		configPath  string
		secretsPath string
	}
}

func NewBrokerApp(configPath, secretsPath string) (app.App, error) {
	a := new(BrokerApp)
	a.config.configPath = configPath
	a.config.secretsPath = secretsPath
	return a, nil
}

func (app *BrokerApp) Run(ctx context.Context) error {
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

	//? WIP
	db.Start()
	correctness.Setup()
	ethereum.Start()
	uniswap.Setup()
	logs.Setup()

	// install GraphQL handlers
	err = gql.InstallHandlers()
	if err != nil {
		return err
	}

	return net.StartServer()
}
