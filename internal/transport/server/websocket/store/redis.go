package store

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client   *redis.Client
	fallback ClientRepository
}

func (r RedisStore) GetAll() []model.Signer {
	return r.fallback.GetAll()
}

func (r RedisStore) GetByPublicKey(publicKey [96]byte) (*websocket.Conn, bool) {
	return r.fallback.GetByPublicKey(publicKey)
}

func (r RedisStore) Set(conn *websocket.Conn, signer model.Signer) {
	r.fallback.Set(conn, signer)
}

func (r RedisStore) Remove(conn *websocket.Conn) {
	r.fallback.Remove(conn)
}

func (r RedisStore) Get(conn *websocket.Conn) (model.Signer, bool) {
	return r.fallback.Get(conn)
}

func NewRedisStore(db database.IRedisDatabase, fallback ClientRepository) ClientRepository {
	return &RedisStore{
		client:   db.GetConnection(),
		fallback: fallback,
	}
}
