package handler

import (
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
)

type Handler struct {
	middleware       middleware.Middleware
	clientRepository store.ClientRepository
}

func NewHandler(middleware middleware.Middleware, clientRepository store.ClientRepository) *Handler {
	return &Handler{
		middleware:       middleware,
		clientRepository: clientRepository,
	}
}
