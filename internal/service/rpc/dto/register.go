package dto

import sia "github.com/pouya-eghbali/go-sia/v2/pkg"

// RegisterFunction is a DTO for registering a function.
type RegisterWorker struct {
	Plugins []Plugin `json:"plugins"`
	CPU     int      `json:"cpu"`
	GPU     int      `json:"gpu"`
	RAM     int      `json:"ram"`
}

func (t *RegisterWorker) Sia() sia.Sia {
	plugins := sia.NewSiaArray[Plugin]().
		AddArray16(t.Plugins, func(s *sia.ArraySia[Plugin], v Plugin) {
			s.EmbedBytes(v.Sia().Bytes())
		})

	return sia.New().
		EmbedBytes(plugins.Bytes()).
		AddInt64(int64(t.CPU)).
		AddInt64(int64(t.GPU)).
		AddInt64(int64(t.RAM))
}

func (t *RegisterWorker) FromSiaBytes(bytes []byte) *RegisterWorker {
	sp := sia.NewArrayFromBytes[Plugin](bytes)

	t.Plugins = sp.ReadArray16(func(s *sia.ArraySia[Plugin]) Plugin {
		p := Plugin{}
		return *p.FromSia(s)
	})

	t.CPU = int(sp.ReadInt64())
	t.GPU = int(sp.ReadInt64())
	t.RAM = int(sp.ReadInt64())

	return t
}
