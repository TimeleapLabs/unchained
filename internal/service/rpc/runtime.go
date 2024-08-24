package rpc

import (
	"net"
)

// Runtime is a type that holds the runtime of a function.
type Runtime string

const (
	Mock Runtime = "Mock"
	Unix Runtime = "Unix"
)

func WithMockTask(name string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Mock,
		}
	}
}

func WithUnixSocket(name string, path string) func(s *Worker) {
	return func(s *Worker) {
		meta := meta{
			runtime: Unix,
			path:    path,
		}

		var err error
		meta.conn, err = net.Dial("unix", path)
		if err != nil {
			panic(err)
		}

		s.functions[name] = meta
	}
}
