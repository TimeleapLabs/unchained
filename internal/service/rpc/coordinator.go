package rpc

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
)

type RemoteWorker struct {
	Worker
	Conn *websocket.Conn
}

type Task struct {
	Worker *websocket.Conn
	Client *websocket.Conn
	CPU    int
	GPU    int
}

// Coordinator is a struct that holds the tasks and workers.
type Coordinator struct {
	Tasks   map[uuid.UUID]Task
	Workers []RemoteWorker
}

// RegisterTask will register a task which a connection provide.
func (r *Coordinator) RegisterTask(taskID uuid.UUID, worker *websocket.Conn, client *websocket.Conn, cpu int, gpu int) {
	r.Tasks[taskID] = Task{
		Worker: worker,
		Client: client,
		CPU:    cpu,
		GPU:    gpu,
	}
}

func (r *Coordinator) GetWorker(conn *websocket.Conn) *RemoteWorker {
	for _, worker := range r.Workers {
		if worker.Conn == conn {
			return &worker
		}
	}

	return nil
}

// UnregisterTask will unregister a task which a connection provide.
func (r *Coordinator) UnregisterTask(taskID uuid.UUID) {
	task, ok := r.Tasks[taskID]
	if !ok {
		return
	}

	worker := r.GetWorker(task.Worker)
	delete(r.Tasks, taskID)

	if worker != nil {
		worker.CPUUsage -= task.CPU
		worker.GPUUsage -= task.GPU
	}
}

// GetTask will return a task which a connection provide.
func (r *Coordinator) GetTask(taskID uuid.UUID) (Task, bool) {
	task, ok := r.Tasks[taskID]
	return task, ok
}

// RegisterWorker will register a worker which a connection provide.
func (r *Coordinator) RegisterWorker(worker *dto.RegisterWorker, conn *websocket.Conn) {
	pluginsMap := make(map[string]dto.Plugin)
	for _, plugin := range worker.Plugins {
		pluginsMap[plugin.Name] = plugin
	}

	r.Workers = append(r.Workers, RemoteWorker{
		Worker: Worker{
			MaxCPU:  worker.CPU,
			MaxGPU:  worker.GPU,
			Plugins: pluginsMap,
		},
		Conn: conn,
	})
}

// UnregisterWorker will unregister a worker which a connection provide.
func (r *Coordinator) UnregisterWorker(conn *websocket.Conn) {
	for i, worker := range r.Workers {
		if worker.Conn == conn {
			r.Workers = append(r.Workers[:i], r.Workers[i+1:]...)
			break
		}
	}
}

// GetWorkers will return a list of workers which provide a function.
func (r *Coordinator) GetWorkers(plugin string, function string) []*RemoteWorker {
	workers := make([]*RemoteWorker, 0, len(r.Workers))

	for _, worker := range r.Workers {
		if _, ok := store.Signers.Load(worker.Conn); !ok {
			r.UnregisterWorker(worker.Conn)
			continue
		}

		if p, ok := worker.Plugins[plugin]; ok {
			if f, ok := p.Functions[function]; ok {
				// Check if the worker has enough resources
				if worker.CPUUsage+f.CPU <= worker.MaxCPU && worker.GPUUsage+f.GPU <= worker.MaxGPU {
					workers = append(workers, &worker)
				}
			}
		}
	}

	return workers
}

// GetRandomWorker will return a random worker which provide a function.
func (r *Coordinator) GetRandomWorker(plugin string, method string) (*RemoteWorker, *config.Function) {
	workers := r.GetWorkers(plugin, method)

	if len(workers) == 0 {
		return nil, nil
	}

	random := rand.Intn(len(workers))
	worker := workers[random]
	function, ok := worker.Plugins[plugin].Functions[method]
	if !ok {
		return nil, nil
	}

	return worker, &function
}

// NewCoordinator creates a new Coordinator.
func NewCoordinator() *Coordinator {
	return &Coordinator{
		Tasks:   make(map[uuid.UUID]Task),
		Workers: make([]RemoteWorker, 0),
	}
}
