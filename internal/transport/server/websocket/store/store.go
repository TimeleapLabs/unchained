package store

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"time"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

var Challenges = xsync.NewMapOf[*websocket.Conn, model.ChallengePacket]()
var OnlineFrostParties = xsync.NewMapOf[string, time.Time]()

type ClientRepository interface {
	Set(conn *websocket.Conn, signer model.Signer)
	Remove(conn *websocket.Conn)
	Get(conn *websocket.Conn) (model.Signer, bool)
	GetAll() []model.Signer
	GetByPublicKey(publicKey [96]byte) (*websocket.Conn, bool)
}
