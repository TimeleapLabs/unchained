package rpc

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
)

func WithMockTask(name string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Mock,
		}
	}
}

func WithWebhookTask(name string, path string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Webhook,
			path:    path,
		}

		packet := dto.RegisterFunction{
			Function: name,
			Runtime:  string(Webhook),
		}
		conn.Send(consts.OpCodeRegisterRpcFunction, packet.Sia().Bytes())
	}
}

func WithWasmTask(name string, path string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Wasm,
			path:    path,
		}

		packet := dto.RegisterFunction{
			Function: name,
			Runtime:  string(Wasm),
		}
		conn.Send(consts.OpCodeRegisterRpcFunction, packet.Sia().Bytes())
	}
}

func WithPythonTask(name string, path string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Python,
			path:    path,
		}

		packet := dto.RegisterFunction{
			Function: name,
			Runtime:  string(Python),
		}
		conn.Send(consts.OpCodeRegisterRpcFunction, packet.Sia().Bytes())
	}
}

func WithDockerTask(name string, path string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Docker,
			path:    path,
		}

		packet := dto.RegisterFunction{
			Function: name,
			Runtime:  string(Docker),
		}
		conn.Send(consts.OpCodeRegisterRpcFunction, packet.Sia().Bytes())
	}
}
