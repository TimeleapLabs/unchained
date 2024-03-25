package clientidentity

type Option func(*clientIdentity) error

func OptionWithName(v string) Option {
	return func(ci *clientIdentity) error {
		ci.config.name = v
		return nil
	}
}

func OptionWithSecretKey(v string) Option {
	return func(ci *clientIdentity) error {
		ci.config.secretKey = v
		return nil
	}
}

func OptionWithEvmWallet(v string) Option {
	return func(ci *clientIdentity) error {
		ci.config.evmWallet = v
		return nil
	}
}
