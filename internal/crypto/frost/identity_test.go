package frost

import (
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bytemare/frost"
	"github.com/bytemare/frost/dkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
)

const (
	numOfSigners    = 4
	minNumOfSigners = 2
)

var (
	testData = []byte("HELLO hello")

	signers = []int{}
)

type FrostIdentityTestSuite struct {
	suite.Suite
	parties []*DistributedSigner
}

func (s *FrostIdentityTestSuite) SetupTest() {
	utils.SetupLogger("info")

	for i := 0; i < numOfSigners; i++ {
		signers = append(signers, rand.Intn(9999))
	}

	r1Messages := []*dkg.Round1Data{}
	for i := 0; i < numOfSigners; i++ {
		r1Msg, party := NewIdentity(signers[i], numOfSigners, minNumOfSigners)
		s.parties = append(s.parties, party)
		r1Messages = append(r1Messages, r1Msg)
	}

	r2Messages := []*dkg.Round2Data{}
	for i := 0; i < numOfSigners; i++ {
		for j := 0; j < numOfSigners; j++ {
			r2Msgs, err := s.parties[i].Update(r1Messages[j])
			assert.NoError(s.T(), err)

			r2Messages = append(r2Messages, r2Msgs...)
		}
	}

	for k := 0; k < numOfSigners; k++ {
		for _, msg := range r2Messages {
			err := s.parties[k].Finalize(msg)
			assert.NoError(s.T(), err)
		}
	}
}

func (s *FrostIdentityTestSuite) TestSign() {
	signers := make([]*MessageSigner, 0, minNumOfSigners)
	commits := make([]*frost.Commitment, 0, minNumOfSigners)

	for i := 0; i < minNumOfSigners; i++ {
		signer, msg, err := s.parties[i].NewSigner(testData)
		assert.NoError(s.T(), err)

		commits = append(commits, msg)
		signers = append(signers, signer)
	}

	for i := 0; i < minNumOfSigners; i++ {
		for j := 0; j < minNumOfSigners; j++ {
			err := signers[i].Confirm(commits[j])
			assert.NoError(s.T(), err)
		}
	}
}

func TestFrostIdentitySuite(t *testing.T) {
	suite.Run(t, new(FrostIdentityTestSuite))
}
