package config

import (
	"time"
)

// System struct hold the internal system configuration.
type System struct {
	Name string `env:"SYSTEM_NAME" env-default:"Unchained" yaml:"name"`
	Log  string `env:"SYSTEM_LOG"  env-default:"info"      yaml:"log"`

	ConfigPath           string
	SecretsPath          string
	AllowGenerateSecrets bool
	ContextPath          string
	Home                 string
	PrintVersion         bool
}

// RPC struct hold the rpc configuration of the application.
type RPC struct {
	Name  string   `yaml:"name"`
	Nodes []string `yaml:"nodes"`
}

// Plugins struct hold the plugins configurations of the application.
type Plugin struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Endpoint  string   `json:"endpoint"`
	Functions []string `json:"functions"`
}

// ProofOfStake struct hold the proof of stake contract's configurations.
type ProofOfStake struct {
	Chain   string `env:"POS_CHAIN"   env-default:"arbitrumSepolia"                            yaml:"chain"`
	Address string `env:"POS_ADDRESS" env-default:"0x965e364987356785b7E89e2Fe7B70f5E5107332d" yaml:"address"`
	Base    int64  `env:"POS_BASE"    env-default:"1"                                          yaml:"base"`
}

// Network struct hold the network configuration of the application.
type Network struct {
	Bind              string        `env:"BIND"               env-default:"0.0.0.0:9123"                    yaml:"bind"`
	CertFile          string        `env:"CERT_FILE"          env-default:""                                yaml:"certFile"`
	KeyFile           string        `env:"KEY_FILE"           env-default:""                                yaml:"keyFile"`
	BrokerURI         string        `env:"BROKER_URI"         env-default:"wss://shinobi.brokers.kenshi.io" yaml:"brokerUri"`
	SubscribedChannel string        `env:"SUBSCRIBED_CHANNEL" env-default:"unchained:"                      yaml:"subscribedChannel"`
	BrokerTimeout     time.Duration `env:"BROKER_TIMEOUT"     env-default:"3s"                              yaml:"brokerTimeout"`
}

type Mongo struct {
	URL      string `env:"Mongo_URL"      yaml:"url"`
	Database string `env:"Mongo_Database" yaml:"database"`
}

// Postgres struct hold the postgres configuration of the application.
type Postgres struct {
	URL string `env:"DATABASE_URL" yaml:"url"`
}

// Secret struct hold the secret keys of the application and loaded from secret.yaml.
type Secret struct {
	Address       string `env:"ADDRESS"         yaml:"address"`
	EvmAddress    string `env:"EVM_ADDRESS"     yaml:"evmAddress"`
	SecretKey     string `env:"SECRET_KEY"      yaml:"secretKey"`
	PublicKey     string `env:"PUBLIC_KEY"      yaml:"publicKey"`
	EvmPrivateKey string `env:"EVM_PRIVATE_KEY" yaml:"evmPrivateKey"`
}

// Function struct hold the function configuration of the application.

// Config struct is the main configuration struct of application.
type Config struct {
	System       System       `yaml:"system"`
	Network      Network      `yaml:"network"`
	RPC          []RPC        `yaml:"rpc"`
	Mongo        Mongo        `yaml:"mongo"`
	Postgres     Postgres     `yaml:"postgres"`
	ProofOfStake ProofOfStake `yaml:"pos"`
	Plugins      []Plugin     `yaml:"plugins"`
	Secret       Secret       `yaml:"secret"`
	Dataframes   []string     `yaml:"dataframes"`
}
