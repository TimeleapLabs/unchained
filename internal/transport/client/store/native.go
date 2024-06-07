package store

import (
	"time"

	"github.com/puzpuzpuz/xsync/v3"
)

type SignerRepository interface {
}

type NativeStore struct {
	signers *xsync.MapOf[string, time.Time]
}

func New() SignerRepository {
	return &NativeStore{
		signers: xsync.NewMapOf[string, time.Time](),
	}
}
