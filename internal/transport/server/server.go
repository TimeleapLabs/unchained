package server

import (
	"fmt"
	"net/http"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/log"
)

func New(options ...func()) {
	for _, option := range options {
		option()
	}

	log.Logger.
		With("Bind", fmt.Sprintf("http://%s", config.App.Network.Bind)).
		Info("Starting a HTTP server")

	server := &http.Server{
		Addr:          config.App.Network.Bind,
		BrokerTimeout: config.App.Network.BrokerTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
