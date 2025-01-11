package rpc

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/worker"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/queue"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
	"golang.org/x/exp/rand"
)

type RemoteWorker struct {
	worker.Worker
	Conn   *websocket.Conn
	Writer *queue.WebSocketWriter
}

type Task struct {
	Worker *websocket.Conn
	Client *queue.WebSocketWriter
	CPU    int
	GPU    int
	RAM    int
}

// Coordinator is a struct that holds the tasks and workers.
type Coordinator struct {
	Tasks   xsync.MapOf[uuid.UUID, Task]
	Workers []RemoteWorker
}

// RegisterTask will register a task which a connection provide.
func (r *Coordinator) RegisterTask(
	taskID uuid.UUID, worker *websocket.Conn,
	client *queue.WebSocketWriter, cpu int, gpu int, ram int) {
	r.Tasks.Store(taskID, Task{
		Worker: worker,
		Client: client,
		CPU:    cpu,
		GPU:    gpu,
		RAM:    ram,
	})

	remoteWorker := r.GetWorker(worker)
	if remoteWorker != nil {
		remoteWorker.CPUUsage += cpu
		remoteWorker.GPUUsage += gpu
		remoteWorker.RAMUsage += ram
	}
}

func (r *Coordinator) GetWorker(conn *websocket.Conn) *RemoteWorker {
	for i := range r.Workers {
		if r.Workers[i].Conn == conn {
			return &r.Workers[i]
		}
	}
	return nil
}

// UnregisterTask will unregister a task which a connection provide.
func (r *Coordinator) UnregisterTask(taskID uuid.UUID) {
	task, ok := r.Tasks.Load(taskID)
	if !ok {
		return
	}

	worker := r.GetWorker(task.Worker)
	r.Tasks.Delete(taskID)

	if worker != nil {
		worker.CPUUsage -= task.CPU
		worker.GPUUsage -= task.GPU
		worker.RAMUsage -= task.RAM
	}
}

// GetTask will return a task which a connection provide.
func (r *Coordinator) GetTask(taskID uuid.UUID) (Task, bool) {
	task, ok := r.Tasks.Load(taskID)
	return task, ok
}

// RegisterWorker will register a worker which a connection provide.
func (r *Coordinator) RegisterWorker(w *dto.RegisterWorker, conn *websocket.Conn) {
	pluginsMap := make(map[string]dto.Plugin)
	for _, plugin := range w.Plugins {
		pluginsMap[plugin.Name] = plugin
	}

	r.Workers = append(r.Workers, RemoteWorker{
		Worker: worker.Worker{
			MaxCPU:  w.CPU,
			MaxGPU:  w.GPU,
			MaxRAM:  w.RAM,
			Plugins: pluginsMap,
		},
		Conn:   conn,
		Writer: queue.NewWebSocketWriter(conn, 10),
	})
}

// UnregisterWorker will unregister a worker which a connection provide.
func (r *Coordinator) UnregisterWorker(conn *websocket.Conn) {
	for i := range r.Workers {
		worker := &r.Workers[i]
		if worker.Conn == conn {
			r.Workers = append(r.Workers[:i], r.Workers[i+1:]...)
			break
		}
	}
}

// GetWorkers will return a list of workers which provide a function.
func (r *Coordinator) GetWorkers(plugin string, function string) []*RemoteWorker {
	workers := make([]*RemoteWorker, 0, len(r.Workers))

	for i := range r.Workers {
		worker := &r.Workers[i] // Get a pointer to the actual slice element

		if _, ok := store.Signers.Load(worker.Conn); !ok {
			r.UnregisterWorker(worker.Conn)
			continue
		}

		if p, ok := worker.Plugins[plugin]; ok {
			if f, ok := p.Functions[function]; ok {
				// Check if the worker has enough resources
				if worker.CPUUsage+f.CPU <= worker.MaxCPU &&
					worker.GPUUsage+f.GPU <= worker.MaxGPU &&
					worker.RAMUsage+f.RAM <= worker.MaxRAM {
					workers = append(workers, worker)
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
		Tasks:   *xsync.NewMapOf[uuid.UUID, Task](),
		Workers: make([]RemoteWorker, 0),
	}
}
