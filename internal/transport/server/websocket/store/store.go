package store

import (
	"github.com/KenshiTech/unchained/internal/crypto/kosk"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

var Consumers = xsync.NewMapOf[*websocket.Conn, bool]()

var Challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
var Signers = xsync.NewMapOf[*websocket.Conn, model.Signer]()
