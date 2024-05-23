package frost

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	numOfSigners    = 10
	minNumOfSigners = 6
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
		signers = append(signers, i+1)
	}

	r1Messages := [][]byte{}
	for i := 0; i < numOfSigners; i++ {
		r1Msg, party := NewIdentity(signers[i], numOfSigners, minNumOfSigners)
		s.parties = append(s.parties, party)
		r1Messages = append(r1Messages, r1Msg)
	}

	r2Messages := [][]byte{}
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
	commits := [][]byte{}

	for i := 0; i < minNumOfSigners; i++ {
		signer, msg, err := s.parties[i].NewSigner(testData)
		assert.NoError(s.T(), err)

		commits = append(commits, msg)
		signers = append(signers, signer)
	}

	signatureShares := make([][]byte, 0, minNumOfSigners)

	for _, signer := range signers {
		for _, commit := range commits {
			signature, err := signer.Confirm(commit)
			assert.NoError(s.T(), err)

			if signature != nil {
				signatureShares = append(signatureShares, signature)
			}
		}
	}

	signature, err := signers[0].Aggregate(signatureShares)
	assert.NoError(s.T(), err)

	ok, err := signers[0].Verify(signature)
	assert.NoError(s.T(), err)
	assert.True(s.T(), ok)
}

func TestFrostIdentitySuite(t *testing.T) {
	suite.Run(t, new(FrostIdentityTestSuite))
}
