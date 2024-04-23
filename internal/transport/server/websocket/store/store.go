package store

import (
	"sync"

	"github.com/TimeleapLabs/unchained/internal/crypto/kosk"
	"github.com/TimeleapLabs/unchained/internal/datasets"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
)

var Consumers = xsync.NewMapOf[*websocket.Conn, bool]()
var BroadcastMutex = xsync.NewMapOf[*websocket.Conn, *sync.Mutex]()

var Challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
var Signers = xsync.NewMapOf[*websocket.Conn, datasets.Signer]()
