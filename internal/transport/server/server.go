package server

import (
	"fmt"
	"net/http"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/log"
)

func New(options ...func()) {
	for _, o := range options {
		o()
	}

	log.Logger.With("Bind", fmt.Sprintf("http://%s", config.App.Broker.Bind)).Info("Server is starting")

	server := &http.Server{
		Addr:              config.App.Broker.Bind,
		ReadHeaderTimeout: config.App.Broker.ReadHeaderTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
