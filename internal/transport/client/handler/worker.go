package handler

import "github.com/TimeleapLabs/unchained/internal/service/frost"

type worker struct {
	frostService frost.Service
}

func NewWorkerHandler(frostService frost.Service) Handler {
	return &worker{
		frostService: frostService,
	}
}
