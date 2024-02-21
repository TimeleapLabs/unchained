package repository

import (
	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

// TODO: Do consumers need KOSK?
var Consumers *xsync.MapOf[*websocket.Conn, bool]

func init() {
	Consumers = xsync.NewMapOf[*websocket.Conn, bool]()
}
