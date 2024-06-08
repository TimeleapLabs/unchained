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
	PrintVersion         bool
}

// RPC struct represent a RPC configs.
type RPC struct {
	Name  string   `yaml:"name"`
	Nodes []string `yaml:"nodes"`
}

// Uniswap struct represent all task's detail of its plugin.
type Uniswap struct {
	Schedule map[string]time.Duration `yaml:"schedule"`
	Tokens   []Token                  `yaml:"tokens"`
}

// EthLog struct represent all task's detail of its plugin.
type EthLog struct {
	Schedule map[string]time.Duration `yaml:"schedule"`
	Events   []Event                  `yaml:"events"`
}

// Frost struct represent all task detail of its plugin.
type Frost struct {
	Schedule  time.Duration `yaml:"schedule"`
	Heartbeat time.Duration `yaml:"heartbeat"`
	Session   string        `yaml:"session"`
}

// Plugins struct holds all applications plugin configs.
type Plugins struct {
	EthLog      *EthLog  `yaml:"logs"`
	Uniswap     *Uniswap `yaml:"uniswap"`
	Frost       *Frost   `yaml:"frost"`
	Correctness []string `yaml:"correctness"`
}

// Event struct represent all events in EthLog plugin.
type Event struct {
	Name          string  `yaml:"name"`
	Chain         string  `yaml:"chain"`
	Abi           string  `yaml:"abi"`
	Event         string  `yaml:"event"`
	Address       string  `yaml:"address"`
	From          *uint64 `yaml:"from"`
	Step          uint64  `yaml:"step"`
	Confirmations uint64  `yaml:"confirmations"`
	Store         bool    `yaml:"store"`
	Send          bool    `yaml:"send"`
}

// Token struct represent all info about a token in plugin.
type Token struct {
	Name   string `yaml:"name"`
	Pair   string `yaml:"pair"`
	Chain  string `yaml:"chain"`
	Delta  int64  `yaml:"delta"`
	Invert bool   `yaml:"invert"`
	Unit   string `yaml:"unit"`
	Send   bool   `yaml:"send"`
	Store  bool   `yaml:"store"`
}

// ProofOfStake struct holds information about POS contract of application.
type ProofOfStake struct {
	Chain   string `env:"POS_CHAIN"   env-default:"arbitrumSepolia"                            yaml:"chain"`
	Address string `env:"POS_ADDRESS" env-default:"0x54550AAfe0df642fbcAde11174250542D0d5FE54" yaml:"address"`
	Base    int64  `env:"POS_BASE"    env-default:"1"                                          yaml:"base"`
}

// Network struct holds all application network configuration.
type Network struct {
	Bind              string        `env:"BIND"               env-default:"0.0.0.0:9123"                    yaml:"bind"`
	BrokerURI         string        `env:"BROKER_URI"         env-default:"wss://shinobi.brokers.kenshi.io" yaml:"brokerUri"`
	SubscribedChannel string        `env:"SUBSCRIBED_CHANNEL" env-default:"unchained:"                      yaml:"subscribedChannel"`
	BrokerTimeout     time.Duration `env:"BROKER_TIMEOUT"     env-default:"3s"                              yaml:"brokerTimeout"`
}

// Postgres struct holds all configs to connect to a pg instance.
type Postgres struct {
	URL string `env:"DATABASE_URL" yaml:"url"`
}

// Redis struct holds all configs to connect to a redis instance.
type Redis struct {
	Dsn string `env:"REDIS_DSN" yaml:"dsn"`
}

// Secret struct hold the secret keys of the application and loaded from secret.yaml.
type Secret struct {
	Address       string `env:"ADDRESS"         yaml:"address"`
	EvmAddress    string `env:"EVM_ADDRESS"     yaml:"evmAddress"`
	SecretKey     string `env:"SECRET_KEY"      yaml:"secretKey"`
	PublicKey     string `env:"PUBLIC_KEY"      yaml:"publicKey"`
	EvmPrivateKey string `env:"EVM_PRIVATE_KEY" yaml:"evmPrivateKey"`

	ShortPublicKey [48]byte `yaml:"-"`
}

// Config struct is the main configuration struct of application.
type Config struct {
	System       System       `yaml:"system"`
	Network      Network      `yaml:"network"`
	RPC          []RPC        `yaml:"rpc"`
	Postgres     Postgres     `yaml:"postgres"`
	Redis        Redis        `yaml:"redis"`
	ProofOfStake ProofOfStake `yaml:"pos"`
	Plugins      Plugins      `yaml:"plugins"`
	Secret       Secret       `yaml:"secret"`
}
