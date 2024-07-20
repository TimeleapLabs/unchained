package runtime

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/tetratelabs/wazero"
	"log"
	"os"
)

func RunWasmFromFile(ctx context.Context, path string, params ...uint64) ([]uint64, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		utils.Logger.With("err", err).Error("Can't open wasm file")
		return nil, err
	}

	return RunWasm(ctx, file, params...)
}

func RunWasm(ctx context.Context, source []byte, params ...uint64) ([]uint64, error) {
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, source)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	// Call the `add` function and print the results to the console.
	add := mod.ExportedFunction("add")
	results, err := add.Call(ctx, params...)
	if err != nil {
		log.Panicf("failed to call add: %v", err)
	}

	return results, nil
}
