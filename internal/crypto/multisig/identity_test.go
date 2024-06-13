package multisig

import (
	"fmt"
	"sync"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

const (
	numOfSigners    = 10
	minNumOfSigners = 6
)

var (
	sessionID = "sessionID"
	testData  = "hello world!"
)

type MultiSigIdentityTestSuite struct {
	suite.Suite
	parties []*DistributedSigner
}

func (s *MultiSigIdentityTestSuite) SetupTest() {
	utils.SetupLogger("info")

	signersIDs := []string{}
	for i := 0; i < numOfSigners; i++ {
		signersIDs = append(signersIDs, fmt.Sprintf("signer-%d", i))
	}

	channels := []<-chan *protocol.Message{}
	for i := 0; i < numOfSigners; i++ {
		party, msgCh := NewIdentity(sessionID, fmt.Sprintf("signer-%d", i), signersIDs, minNumOfSigners)
		s.parties = append(s.parties, party)
		channels = append(channels, msgCh)
	}

	wg := sync.WaitGroup{}
	for _, channel := range channels {
		wg.Add(1)
		go func(channel <-chan *protocol.Message) {
			for msg := range channel {
				for _, p := range s.parties {
					if msg.Broadcast || msg.IsFor(p.ID) {
						isReady, err := p.Confirm(msg)
						assert.NoError(s.T(), err)

						if isReady {
							wg.Done()
						}
					}
				}
			}
		}(channel)
	}

	wg.Wait()
}

func (s *MultiSigIdentityTestSuite) TestSign() {
	signers := []*MessageSigner{}
	channels := []<-chan *protocol.Message{}

	for _, party := range s.parties {
		signer, msgCh, err := party.NewSigner([]byte(testData))
		assert.NoError(s.T(), err)

		signers = append(signers, signer)
		channels = append(channels, msgCh)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	for _, channel := range channels {
		go func(channel <-chan *protocol.Message) {
			for msg := range channel {
				for _, p := range signers {
					if msg.Broadcast || msg.IsFor(p.ID) {
						isReady, err := p.Confirm(msg)
						assert.NoError(s.T(), err)

						if isReady {
							wg.Done()
						}
					}
				}
			}
		}(channel)
	}

	wg.Wait()
}

func TestMultiSigIdentitySuite(t *testing.T) {
	suite.Run(t, new(MultiSigIdentityTestSuite))
}
