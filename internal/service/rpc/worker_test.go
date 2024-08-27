package rpc

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	UnixSocketPath = "/tmp/test.sock"
)

var (
	SamplePacket = dto.NewRequest("test", nil, [48]byte{}, "txHash")
)

type WorkerTestSuite struct {
	suite.Suite
	service *Worker
	server  *net.UnixListener
}

func handleConnection(t *testing.T, conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		assert.NoError(t, err)

		t.Log("Received: ", buf[0:n])
		_, err = conn.Write(buf[0:n])
		assert.NoError(t, err)
	}
}

func (s *WorkerTestSuite) SetupTest() {
	utils.SetupLogger("info")

	_ = os.Remove(UnixSocketPath)

	var err error
	s.server, err = net.ListenUnix("unix", &net.UnixAddr{Name: UnixSocketPath, Net: "unix"})
	s.Require().NoError(err)

	go func() {
		for {
			conn, err := s.server.Accept()
			s.Require().NoError(err)

			go handleConnection(s.T(), conn)
		}
	}()

	s.service = NewWorker(
		WithMockTask("test"),
		WithUnixSocket("unix-test", UnixSocketPath),
	)
}

func (s *WorkerTestSuite) TestRunFunction() {
	s.Run("Should run successfully", func() {
		result, err := s.service.RunFunction(context.TODO(), "test", &SamplePacket)
		s.NoError(err)
		s.Equal(SamplePacket.Sia().Bytes(), result)
	})

	s.Run("Run non-existing func, Should return err", func() {
		_, err := s.service.RunFunction(context.TODO(), "non-existing-test", &SamplePacket)
		s.Error(err, consts.ErrInternalError)
	})

	s.Run("Should run successfully", func() {
		result, err := s.service.RunFunction(context.TODO(), "unix-test", &SamplePacket)
		s.NoError(err)

		gotPacket := new(dto.RPCRequest).FromSiaBytes(result)
		s.Equal(SamplePacket, *gotPacket)
	})
}

func (s *WorkerTestSuite) TearDownTest() {
	err := s.server.Close()
	s.Require().NoError(err)

	_ = os.Remove(UnixSocketPath)
}

func TestWorkerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerTestSuite))
}
