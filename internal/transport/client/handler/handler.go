package handler

import (
	"github.com/KenshiTech/unchained/internal/service/correctness"
	"github.com/KenshiTech/unchained/internal/service/evmlog"
	"github.com/KenshiTech/unchained/internal/service/uniswap"
)

type Handler struct {
	correctness *correctness.Service
	uniswap     *uniswap.Service
	evmlog      *evmlog.Service
}

func New(
	correctness *correctness.Service,
	uniswap *uniswap.Service,
	evmlog *evmlog.Service,
) *Handler {
	return &Handler{
		correctness: correctness,
		uniswap:     uniswap,
		evmlog:      evmlog,
	}
}
