package frost

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/suite"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"sync"
	"testing"
)

var (
	OnlinePeers = []string{
		"0x0000000000000000000000000000000000000001", "0x0000000000000000000000000000000000000002", "0x0000000000000000000000000000000000000003", "0x0000000000000000000000000000000000000004", "0x0000000000000000000000000000000000000005", "0x0000000000000000000000000000000000000006", "0x0000000000000000000000000000000000000007", "0x0000000000000000000000000000000000000008", "0x0000000000000000000000000000000000000009", "0x0000000000000000000000000000000000000010",
	}
	AllPeers = []string{
		"0x0000000000000000000000000000000000000001", "0x0000000000000000000000000000000000000002", "0x0000000000000000000000000000000000000003", "0x0000000000000000000000000000000000000004", "0x0000000000000000000000000000000000000005", "0x0000000000000000000000000000000000000006", "0x0000000000000000000000000000000000000007", "0x0000000000000000000000000000000000000008", "0x0000000000000000000000000000000000000009", "0x0000000000000000000000000000000000000010", "0x0000000000000000000000000000000000000011", "0x0000000000000000000000000000000000000012", "0x0000000000000000000000000000000000000013",
	}
)

type FrostServiceTestSuite struct {
	suite.Suite

	services []Service
}

func (s *FrostServiceTestSuite) SetupTest() {
	utils.SetupLogger("info")

	mockPos := pos.NewMock(AllPeers)

	for i := 0; i < len(OnlinePeers); i++ {
		s.services = append(s.services, New(mockPos))
	}

}

func (s *FrostServiceTestSuite) TestSyncSigners() {
	channels := []<-chan *protocol.Message{}
	for i := 0; i < len(OnlinePeers); i++ {

		config.App.Plugins.Frost = &config.Frost{Session: "session-test"}
		config.App.Secret.EvmAddress = OnlinePeers[i]
		config.App.Secret.PublicKey = "a7f10f30e905b5c7a060f6f761ef8004e09b8faae72a09d7ddbc0c3b75cd98c80bbd79660abaca8242a3ce95e01fdbb51960bf8e67489aa4882b56fd2ed0691c0e9754d0856baeac0424bd38194f595e054eb17068434546dec24d548fb3ec82"

		handshakeChannel, err := s.services[i].SyncSigners(context.TODO(), OnlinePeers)
		s.NoError(err)

		channels = append(channels, handshakeChannel)
	}

	wg := sync.WaitGroup{}
	for i, channel := range channels {
		wg.Add(1)
		go func(channel <-chan *protocol.Message) {
			for msg := range channel {
				for j, _ := range s.services {
					isReady, err := s.services[j].ConfirmHandshakeRaw(context.TODO(), msg)
					s.NoError(err)

					if isReady {
						wg.Done()
					}
				}
			}
		}(channel)
	}

	wg.Wait()
}

func (s *FrostServiceTestSuite) TestSigning() {

}

func TestFrostServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FrostServiceTestSuite))
}
