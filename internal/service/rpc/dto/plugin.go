package dto

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/websocket/queue"
	"github.com/gorilla/websocket"
)

// Runtime is a type that holds the runtime of a function.
type Runtime string

type Plugin struct {
	Name      string
	Conn      *websocket.Conn
	Writer    *queue.WebSocketWriter
	Runtime   Runtime
	Functions map[string]config.Function
}

func (p *Plugin) Sia() sia.Sia {
	functions := make([]config.Function, 0, len(p.Functions))
	for _, v := range p.Functions {
		functions = append(functions, v)
	}

	s := sia.New().
		AddString8(p.Name).
		AddString8(string(p.Runtime))

	sarr := sia.NewArray[config.Function](&s)
	sarr.AddArray16(functions, func(s *sia.ArraySia[config.Function], v config.Function) {
		s.AddString8(v.Name)
		s.AddInt64(int64(v.Timeout))
		s.AddInt64(int64(v.CPU))
		s.AddInt64(int64(v.GPU))
		s.AddInt64(int64(v.RAM))
	})

	return s
}

func (p *Plugin) FromSia(s sia.Sia) *Plugin {
	p.Name = s.ReadString8()
	p.Runtime = Runtime(s.ReadString8())

	sarr := sia.NewArray[config.Function](&s)
	functions := sarr.ReadArray16(func(s *sia.ArraySia[config.Function]) config.Function {
		f := config.Function{}
		f.Name = s.ReadString8()
		f.Timeout = int(s.ReadInt64())
		f.CPU = int(s.ReadInt64())
		f.GPU = int(s.ReadInt64())
		f.RAM = int(s.ReadInt64())
		return f
	})

	p.Functions = make(map[string]config.Function)
	for _, v := range functions {
		p.Functions[v.Name] = v
	}

	return p
}

func (p *Plugin) FromSiaBytes(bytes []byte) *Plugin {
	s := sia.NewFromBytes(bytes)
	return p.FromSia(s)
}
