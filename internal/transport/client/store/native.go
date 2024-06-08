package store

import (
	"time"

	"github.com/puzpuzpuz/xsync/v3"
)

type SignerRepository interface {
	SetSignerIsAlive(evmAddress string)
}

type NativeStore struct {
	signers *xsync.MapOf[string, time.Time]
}

func (n NativeStore) SetSignerIsAlive(evmAddress string) {
	n.signers.Store(evmAddress, time.Now())
}

func New() SignerRepository {
	return &NativeStore{
		signers: xsync.NewMapOf[string, time.Time](),
	}
}
