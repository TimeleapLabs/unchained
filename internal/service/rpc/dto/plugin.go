package dto

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/queue"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
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

	s.EmbedBytes(sia.NewSiaArray[config.Function]().
		AddArray16(functions, func(s *sia.ArraySia[config.Function], v config.Function) {
			s.AddString8(v.Name)
			s.AddInt64(int64(v.CPU))
			s.AddInt64(int64(v.GPU))
			s.AddInt64(int64(v.RAM))
		}).Bytes())

	return s
}

func (p *Plugin) FromSia(s sia.Sia) *Plugin {
	p.Name = s.ReadString8()
	p.Runtime = Runtime(s.ReadString8())

	sa := sia.NewArrayFromBytes[config.Function](s.Bytes()[s.Offset():])
	functions := sa.ReadArray16(func(s *sia.ArraySia[config.Function]) config.Function {
		f := config.Function{}
		f.Name = s.ReadString8()
		f.CPU = int(s.ReadInt64())
		f.GPU = int(s.ReadInt64())
		f.RAM = int(s.ReadInt64())
		return f
	})

	p.Functions = make(map[string]config.Function)
	for _, v := range functions {
		p.Functions[v.Name] = v
	}

	s.Seek(s.Offset() + sa.Offset())
	return p
}

func (p *Plugin) FromSiaBytes(bytes []byte) *Plugin {
	s := sia.NewFromBytes(bytes)
	return p.FromSia(s)
}
