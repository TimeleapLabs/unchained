package store

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

type NativeStore struct {
	signers *xsync.MapOf[*websocket.Conn, model.Signer]
}

func (n *NativeStore) GetByPublicKey(publicKey [96]byte) (*websocket.Conn, bool) {
	var connection *websocket.Conn
	n.signers.Range(func(conn *websocket.Conn, signerInMap model.Signer) bool {
		publicKeyInUse := signerInMap.PublicKey == publicKey
		if publicKeyInUse {
			connection = conn
		}
		return !publicKeyInUse
	})

	return connection, connection != nil
}

func (n *NativeStore) GetAll() []model.Signer {
	signers := []model.Signer{}

	n.signers.Range(func(_ *websocket.Conn, value model.Signer) bool {
		signers = append(signers, value)
		return true
	})

	return signers
}

func (n *NativeStore) Set(conn *websocket.Conn, signer model.Signer) {
	n.signers.Store(conn, signer)
}

func (n *NativeStore) Remove(conn *websocket.Conn) {
	n.signers.Delete(conn)
}

func (n *NativeStore) Get(conn *websocket.Conn) (model.Signer, bool) {
	return n.signers.Load(conn)
}

func New() ClientRepository {
	return &NativeStore{
		signers: xsync.NewMapOf[*websocket.Conn, model.Signer](),
	}
}
