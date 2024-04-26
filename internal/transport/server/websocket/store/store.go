package store

import (
	"github.com/TimeleapLabs/unchained/internal/model"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

var Challenges = xsync.NewMapOf[*websocket.Conn, model.ChallengePacket]()
var Signers = xsync.NewMapOf[*websocket.Conn, model.Signer]()
