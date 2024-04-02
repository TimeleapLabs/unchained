package handler

import (
	"github.com/KenshiTech/unchained/service/evmlog"
	"github.com/KenshiTech/unchained/service/uniswap"
)

type Handler struct {
	uniswap *uniswap.Service
	evmlog  *evmlog.Service
}

func New(
	uniswap *uniswap.Service,
	evmlog *evmlog.Service,
) *Handler {
	return &Handler{
		uniswap: uniswap,
		evmlog:  evmlog,
	}
}
