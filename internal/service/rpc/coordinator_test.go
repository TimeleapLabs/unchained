package rpc

import (
	"testing"

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
	s.service.RegisterWorker("test-worker", conn)
	gotConns := s.service.GetWorkers("test-worker")
	s.Len(gotConns, 1)
	s.Equal(conn, gotConns[0])

	s.service.UnregisterWorker("test-worker", conn)
	gotConns = s.service.GetWorkers("test-worker")
	s.Len(gotConns, 0)
}

func (s *CoordinatorTestSuite) TestCoordinator_RegisterTask() {
	conn := &websocket.Conn{}

	taskID, err := uuid.NewUUID()
	s.NoError(err)

	s.service.RegisterTask(taskID, conn)
	gotConn := s.service.GetTask(taskID)
	s.Equal(conn, gotConn)

	s.service.UnregisterTask(taskID)
	gotConn = s.service.GetTask(taskID)
	s.Nil(gotConn)
}

func TestCoordinatorSuite(t *testing.T) {
	suite.Run(t, new(CoordinatorTestSuite))
}
