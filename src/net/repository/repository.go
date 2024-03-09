package repository

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

// TODO: Do consumers need KOSK?
var Consumers *xsync.MapOf[*websocket.Conn, bool]
var BroadcastMutex *xsync.MapOf[*websocket.Conn, *sync.Mutex]

func init() {
	Consumers = xsync.NewMapOf[*websocket.Conn, bool]()
	BroadcastMutex = xsync.NewMapOf[*websocket.Conn, *sync.Mutex]()
}
