package rpc

import "github.com/google/uuid"

func NewTaskID() (uuid.UUID, error) {
	taskID, err := uuid.NewV7()
	return taskID, err
}
