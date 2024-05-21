package frost

type Service interface {
	PushSigners() error
}

type service struct {
	reserveSigners []bool
	currentSigners []bool
}

func New() Service {
	return &service{}
}
