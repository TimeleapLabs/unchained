package server

import (
	"fmt"
	"net/http"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

// New creates a new HTTP server.
func New(options ...func()) {
	for _, option := range options {
		option()
	}

	server := &http.Server{
		Addr:              config.App.Network.Bind,
		ReadHeaderTimeout: config.App.Network.Broker.Timeout,
	}

	if config.App.Network.CertFile != "" && config.App.Network.KeyFile != "" {
		utils.Logger.
			With("Bind", fmt.Sprintf("https://%s", config.App.Network.Bind)).
			With("CertFile", config.App.Network.CertFile).
			With("KeyFile", config.App.Network.KeyFile).
			Info("Starting a HTTPS server")

		err := server.ListenAndServeTLS(config.App.Network.CertFile, config.App.Network.KeyFile)
		if err != nil {
			panic(err)
		}

		return
	}

	utils.Logger.
		With("Bind", fmt.Sprintf("http://%s", config.App.Network.Bind)).
		Info("Starting a HTTP server")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
