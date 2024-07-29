package runtime

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunWasm(t *testing.T) {
	result, err := RunWasmFromFile(context.TODO(), "./wasm_add_bg.wasm", []byte{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, result[0], 3)
}
