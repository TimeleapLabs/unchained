package store

import (
	"github.com/TimeleapLabs/unchained/internal/crypto/kosk"
	"github.com/TimeleapLabs/unchained/internal/model"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

var Challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
var Signers = xsync.NewMapOf[*websocket.Conn, model.Signer]()
