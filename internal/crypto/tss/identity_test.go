package tss

import (
	"fmt"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/tss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	numOfSigners    = 4
	minNumOfSigners = 2
)

var (
	testData = []byte("HELLO hello")

	signers = []string{}
)

type TssIdentityTestSuite struct {
	suite.Suite
	parties []*DistributedSigner
}

func (s *TssIdentityTestSuite) SetupTest() {
	utils.SetupLogger("info")

	for i := 0; i < numOfSigners; i++ {
		signers = append(signers, fmt.Sprintf("signer-%d", i))
	}

	outCh := make(chan tss.Message, len(signers))
	done := make(chan struct{})

	for i := 0; i < numOfSigners; i++ {
		party := NewIdentity(i, signers, outCh, done, minNumOfSigners)
		s.parties = append(s.parties, party)
	}

	wait := 0
keygen:
	for {
		select {
		case msg := <-outCh:
			dest := msg.GetTo()

			if dest == nil { // broadcast!
				for _, p := range s.parties {
					if p.PartyID.Index == msg.GetFrom().Index {
						continue
					}

					go func(signer *DistributedSigner) {
						err := signer.Update(msg)
						if err != nil {
							utils.Logger.Error(err.Error())
						}
					}(p)
				}
			} else { // point-to-point!
				if dest[0].Index == msg.GetFrom().Index {
					return
				}
				go func() {
					err := s.parties[dest[0].Index].Update(msg)
					if err != nil {
						utils.Logger.Error(err.Error())
					}
				}()
			}

		case <-done:
			wait++
			if wait == numOfSigners {
				break keygen
			}
		}
	}
}

func (s *TssIdentityTestSuite) TestSign() {
	outCh := make(chan tss.Message, minNumOfSigners+1)
	endCh := make(chan common.SignatureData, minNumOfSigners+1)
	msgSigners := make([]*MessageSigner, 0, minNumOfSigners+1)

	for i := 0; i < minNumOfSigners+1; i++ {
		signer, err := s.parties[i].NewSigning(s.parties[:minNumOfSigners+1], testData, outCh, endCh)
		assert.NoError(s.T(), err)

		msgSigners = append(msgSigners, signer)
	}

	wait := 0
signing:
	for {
		select {
		case msg := <-outCh:
			dest := msg.GetTo()
			if dest == nil {
				for _, p := range msgSigners {
					if p.party.PartyID().Index == msg.GetFrom().Index {
						continue
					}

					go func(signer *MessageSigner, msg tss.Message) {
						err := signer.AckSignature(msg)
						if err != nil {
							utils.Logger.Error(err.Error())
						}
					}(p, msg)
				}
			} else {
				if dest[0].Index == msg.GetFrom().Index {
					common.Logger.Fatalf("party %d tried to send a message to itself (%d)", dest[0].Index, msg.GetFrom().Index)
				}

				go func(signer *MessageSigner, msg tss.Message) {
					err := signer.AckSignature(msg)
					if err != nil {
						utils.Logger.Error(err.Error())
					}
				}(msgSigners[dest[0].Index], msg)
			}
		case <-endCh:
			wait++
			if wait == numOfSigners-1 {
				break signing
			}
		}
	}
}

func TestTssIdentitySuite(t *testing.T) {
	suite.Run(t, new(TssIdentityTestSuite))
}
