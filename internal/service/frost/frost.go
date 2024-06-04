package frost

type Service interface {
	InitSigner() error
}

type service struct {
	reserveSigners []bool
	currentSigners []bool
}

func New() Service {
	return &service{}
}
