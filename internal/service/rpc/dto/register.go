package dto

import sia "github.com/TimeleapLabs/go-sia/v2/pkg"

// RegisterFunction is a DTO for registering a function.
type RegisterWorker struct {
	Plugins []Plugin `json:"plugins"`
	CPU     int      `json:"cpu"`
	GPU     int      `json:"gpu"`
	RAM     int      `json:"ram"`
}

func (t *RegisterWorker) Sia() sia.Sia {
	s := sia.New()

	sarr := sia.NewArray[Plugin](&s)
	sarr.
		AddArray16(t.Plugins, func(s *sia.ArraySia[Plugin], v Plugin) {
			s.EmbedBytes(v.Sia().Bytes())
		})

	return s.
		AddInt64(int64(t.CPU)).
		AddInt64(int64(t.GPU)).
		AddInt64(int64(t.RAM))
}

func (t *RegisterWorker) FromSiaBytes(bytes []byte) *RegisterWorker {
	s := sia.NewFromBytes(bytes)
	sarr := sia.NewArray[Plugin](&s)

	t.Plugins = sarr.ReadArray16(func(s *sia.ArraySia[Plugin]) Plugin {
		p := Plugin{}
		return *p.FromSia(s)
	})

	t.CPU = int(s.ReadInt64())
	t.GPU = int(s.ReadInt64())
	t.RAM = int(s.ReadInt64())

	return t
}
