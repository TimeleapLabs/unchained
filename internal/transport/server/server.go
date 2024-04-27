package server

import (
	"fmt"
	"net/http"

	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/config"
)

func New(options ...func()) {
	for _, option := range options {
		option()
	}

	utils.Logger.
		With("Bind", fmt.Sprintf("http://%s", config.App.Network.Bind)).
		Info("Starting a HTTP server")

	server := &http.Server{
		Addr:              config.App.Network.Bind,
		ReadHeaderTimeout: config.App.Network.BrokerTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
