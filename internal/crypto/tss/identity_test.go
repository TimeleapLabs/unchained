package tss

import (
	"fmt"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/tss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	numOfSigners    = 4
	minNumOfSigners = 2
)

var (
	testData       = []byte("HELLO hello")
	secondTestData = []byte("HELLO hello 2")
	signers        = []string{}
)

type TssIdentityTestSuite struct {
	suite.Suite
	parties []*Signer
}

func (s *TssIdentityTestSuite) SetupTest() {
	utils.SetupLogger("info")

	for i := 0; i < numOfSigners; i++ {
		signers = append(signers, fmt.Sprintf("signer-%d", i))
	}

	outCh := make(chan tss.Message, len(signers))
	wg := make(chan struct{})

	for i := 0; i < numOfSigners; i++ {
		party := NewIdentity(i, signers, outCh, wg, minNumOfSigners)
		s.parties = append(s.parties, party)
	}

	wait := 0
keygen:
	for {
		select {
		case msg := <-outCh:
			dest := msg.GetTo()

			if dest == nil { // broadcast!
				for _, P := range s.parties {
					if P.PartyID.Index == msg.GetFrom().Index {
						continue
					}

					go P.Update(msg)
				}
			} else { // point-to-point!
				if dest[0].Index == msg.GetFrom().Index {
					fmt.Printf("party %d tried to send a message to itself (%d)", dest[0].Index, msg.GetFrom().Index)
					return
				}
				go s.parties[dest[0].Index].Update(msg)
			}

		case <-wg:
			wait++
			if wait == numOfSigners {
				break keygen
			}
		}
	}
}

func (s *TssIdentityTestSuite) TestSign() {
	outCh := make(chan tss.Message, len(s.parties))
	endCh := make(chan common.SignatureData, len(s.parties))
	msgSigners := []*MessageSigner{}

	for i := 0; i < numOfSigners; i++ {
		go func() {
			signer, err := s.parties[i].NewSigning(testData, outCh, endCh)
			assert.NoError(s.T(), err)

			msgSigners = append(msgSigners, signer)
		}()
	}

	//	wait := 0
	//signing:
	for {
		select {
		case msg := <-outCh:
			dest := msg.GetTo()
			if dest == nil {
				for _, P := range msgSigners {
					if P.party.PartyID().Index == msg.GetFrom().Index {
						continue
					}

					go P.AckSignature(msg)
				}
			} else {
				if dest[0].Index == msg.GetFrom().Index {
					common.Logger.Fatalf("party %d tried to send a message to itself (%d)", dest[0].Index, msg.GetFrom().Index)
				}

				go msgSigners[dest[0].Index].AckSignature(msg)
			}
		case <-endCh:
			fmt.Println("end")
		}
	}

	//s.Run("Should verify the message correctly", func() {
	//	isVerified, err := s.identity.Verify()
	//	assert.NoError(s.T(), err)
	//	assert.True(s.T(), isVerified)
	//})
}

func TestTssIdentitySuite(t *testing.T) {
	suite.Run(t, new(TssIdentityTestSuite))
}
