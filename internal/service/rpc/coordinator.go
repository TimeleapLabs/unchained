package rpc

import (
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
)

// Coordinator is a struct that holds the tasks and workers.
type Coordinator struct {
	Tasks   map[uuid.UUID]*websocket.Conn
	Workers map[string][]*websocket.Conn
}

// RegisterTask will register a task which a connection provide.
func (r *Coordinator) RegisterTask(taskID uuid.UUID, conn *websocket.Conn) {
	r.Tasks[taskID] = conn
}

// UnregisterTask will unregister a task which a connection provide.
func (r *Coordinator) UnregisterTask(taskID uuid.UUID) {
	delete(r.Tasks, taskID)
}

// GetTask will return a task which a connection provide.
func (r *Coordinator) GetTask(taskID uuid.UUID) *websocket.Conn {
	return r.Tasks[taskID]
}

// RegisterWorker will register a worker which a connection provide.
func (r *Coordinator) RegisterWorker(function string, conn *websocket.Conn) {
	r.Workers[function] = append(r.Workers[function], conn)
}

// UnregisterWorker will unregister a worker which a connection provide.
func (r *Coordinator) UnregisterWorker(function string, conn *websocket.Conn) {
	workers := r.Workers[function]
	for i, c := range workers {
		if c == conn {
			r.Workers[function] = append(workers[:i], workers[i+1:]...)
			break
		}
	}
}

// GetWorkers will return a list of workers which provide a function.
func (r *Coordinator) GetWorkers(function string) []*websocket.Conn {
	return r.Workers[function]
}

// GetRandomWorker will return a random worker which provide a function.
func (r *Coordinator) GetRandomWorker(function string) *websocket.Conn {
	workers := r.Workers[function]
	available := make([]*websocket.Conn, 0, len(workers))

	for _, worker := range workers {
		if _, ok := store.Signers.Load(worker); ok {
			available = append(available, worker)
		}
	}

	if len(available) == 0 {
		return nil
	}

	r.Workers[function] = available
	random := rand.Intn(len(available))

	return available[random]
}

// NewCoordinator creates a new Coordinator.
func NewCoordinator() *Coordinator {
	return &Coordinator{
		Tasks:   make(map[uuid.UUID]*websocket.Conn),
		Workers: make(map[string][]*websocket.Conn),
	}
}
