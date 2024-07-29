package runtime

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/tetratelabs/wazero"
	"log"
	"os"
	"unsafe"
)

func RunWasmFromFile(ctx context.Context, path string, params []byte) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		utils.Logger.With("err", err).Error("Can't open wasm file")
		return nil, err
	}

	return RunWasm(ctx, file, params)
}

func RunWasm(ctx context.Context, source []byte, params []byte) ([]byte, error) {
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	mod, err := r.Instantiate(ctx, source)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	add := mod.ExportedFunction("run")

	memory := mod.Memory()
	
	ptr := uint64(*(*int32)(unsafe.Pointer(&params)))

	results, err := add.Call(ctx, ptr, uint64(len(params)))
	if err != nil {
		log.Panicf("failed to call add: %v", err)
	}

	resultPtr := results[0]
	resultLen := results[1]
	result, ok := memory.Read(uint32(resultPtr), uint32(resultLen))
	if !ok {
		log.Panicf("failed to read result from memory")
	}

	return result, nil
}
