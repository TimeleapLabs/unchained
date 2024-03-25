package xerrors

import (
	"fmt"
)

var (
	ErrNilArgs = func(v interface{}) error { return fmt.Errorf("nil args: %v", v) }
)
