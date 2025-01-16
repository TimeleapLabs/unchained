package store

import (
	"github.com/TimeleapLabs/unchained/internal/model"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

var Signers = xsync.NewMapOf[*websocket.Conn, model.Signer]()
