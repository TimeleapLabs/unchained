package evmlog

import (
	"github.com/TimeleapLabs/unchained/internal/model"
)

func sortEventArgs(lhs model.EventLogArg, rhs model.EventLogArg) int {
	if lhs.Name < rhs.Name {
		return -1
	}
	return 1
}
