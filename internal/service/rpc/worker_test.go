package rpc

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WorkerTestSuite struct {
	suite.Suite
	service *Worker
}

func (s *WorkerTestSuite) SetupTest() {
	utils.SetupLogger("info")

	s.service = NewWorker(
		WithMockTask("test"),
	)
}

func (s *WorkerTestSuite) TestRunFunction() {
	s.Run("Should run successfully", func() {
		result, err := s.service.RunFunction(context.TODO(), "test", []byte("hello world"))
		s.NoError(err)
		s.Equal("hello world", string(result))
	})

	s.Run("Run non-existing func, Should return err", func() {
		_, err := s.service.RunFunction(context.TODO(), "non-existing-test", []byte("hello world"))
		s.Error(err, consts.ErrInternalError)
	})
}

func TestWorkerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerTestSuite))
}
