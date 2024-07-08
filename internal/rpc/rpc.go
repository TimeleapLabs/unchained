package rpc

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
)

type RPC struct {
	Tasks   map[uuid.UUID]*websocket.Conn
	Workers map[string][]*websocket.Conn
}

func New() *RPC {
	return &RPC{
		Tasks:   make(map[uuid.UUID]*websocket.Conn),
		Workers: make(map[string][]*websocket.Conn),
	}
}

func (r *RPC) RegisterTask(taskID uuid.UUID, conn *websocket.Conn) {
	r.Tasks[taskID] = conn
}

func (r *RPC) UnregisterTask(taskID uuid.UUID) {
	delete(r.Tasks, taskID)
}

func (r *RPC) GetTask(taskID uuid.UUID) *websocket.Conn {
	return r.Tasks[taskID]
}

func (r *RPC) NewTaskID() (uuid.UUID, error) {
	taskID, err := uuid.NewV7()
	return taskID, err
}

func (r *RPC) RegisterWorker(function string, conn *websocket.Conn) {
	r.Workers[function] = append(r.Workers[function], conn)
}

func (r *RPC) UnregisterWorker(function string, conn *websocket.Conn) {
	workers := r.Workers[function]
	for i, c := range workers {
		if c == conn {
			r.Workers[function] = append(workers[:i], workers[i+1:]...)
			break
		}
	}
}

func (r *RPC) GetWorkers(function string) []*websocket.Conn {
	return r.Workers[function]
}

func (r *RPC) GetRandomWorker(function string) *websocket.Conn {
	workers := r.Workers[function]
	if len(workers) == 0 {
		return nil
	}

	random := rand.Intn(len(workers))
	return workers[random]
}