package dto

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/google/uuid"
)

// WorkerOverload is a DTO for worker overload.
type WorkerOverload struct {
	FailedTaskID uuid.UUID `json:"failed_task_id"`
	CPU          int       `json:"cpu"`
	GPU          int       `json:"gpu"`
	RAM          int       `json:"ram"`
}

func (t *WorkerOverload) Sia() sia.Sia {
	uuidBytes, err := t.FailedTaskID.MarshalBinary()

	if err != nil {
		panic(err)
	}

	return sia.New().
		AddByteArray8(uuidBytes).
		AddInt64(int64(t.CPU)).
		AddInt64(int64(t.GPU)).
		AddInt64(int64(t.RAM))
}

func (t *WorkerOverload) FromSiaBytes(bytes []byte) *WorkerOverload {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.FailedTaskID.UnmarshalBinary(uuidBytes)

	if err != nil {
		panic(err)
	}

	t.CPU = int(s.ReadInt64())
	t.GPU = int(s.ReadInt64())
	t.RAM = int(s.ReadInt64())

	return t
}
