package rpc

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/queue"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/suite"
)

type CoordinatorTestSuite struct {
	suite.Suite
	service *Coordinator
}

func (s *CoordinatorTestSuite) SetupTest() {
	utils.SetupLogger("info")

	s.service = NewCoordinator()
}

func (s *CoordinatorTestSuite) TestCoordinator_RegisterWorker() {
	conn := &websocket.Conn{}
	worker := dto.RegisterWorker{
		CPU: 100,
		GPU: 1,
		Plugins: []dto.Plugin{
			{
				Name: "test-plugin",
				Functions: map[string]config.Function{
					"test-function": {
						Name: "test-function",
						CPU:  10,
					},
				},
			},
		},
	}
	s.service.RegisterWorker(&worker, conn)
	store.Signers.Store(conn, model.Signer{ID: 0})

	gotConns := s.service.GetWorkers("test-plugin", "test-function")
	s.Len(gotConns, 1)
	s.Equal(conn, gotConns[0].Conn)

	s.service.UnregisterWorker(conn)
	gotConns = s.service.GetWorkers("test-plugin", "test-function")
	s.Len(gotConns, 0)
}

func (s *CoordinatorTestSuite) TestCoordinator_RegisterTask() {
	worker := &websocket.Conn{}
	client := &websocket.Conn{}
	writer := queue.NewWebSocketWriter(client, 10)

	taskID, err := uuid.NewUUID()
	s.NoError(err)

	s.service.RegisterTask(taskID, worker, writer, 100, 1)
	task, _ := s.service.GetTask(taskID)
	s.Equal(worker, task.Worker)
	s.Equal(writer, task.Client)

	s.service.UnregisterTask(taskID)
	task, _ = s.service.GetTask(taskID)
	s.Nil(task.Worker)
	s.Nil(task.Client)
}

func TestCoordinatorSuite(t *testing.T) {
	suite.Run(t, new(CoordinatorTestSuite))
}
