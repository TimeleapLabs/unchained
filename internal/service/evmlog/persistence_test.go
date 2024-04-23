package evmlog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	TestKey          = "KEY"
	TestValue uint64 = 100
)

type PersistenceTestSuite struct {
	suite.Suite
	badgerRepository *Badger
}

func (s *PersistenceTestSuite) SetupTest() {
	s.badgerRepository = NewBadger("./context")
}

func (s *PersistenceTestSuite) TearDownSuite() {
	err := os.RemoveAll("./context")
	assert.NoError(s.T(), err)
}

func (s *PersistenceTestSuite) TestWriteUint64() {
	err := s.badgerRepository.WriteUint64(TestKey, TestValue)
	assert.NoError(s.T(), err)

	gotValue, err := s.badgerRepository.ReadUInt64(TestKey)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), TestValue, gotValue)
}

func TestPersistenceSuite(t *testing.T) {
	suite.Run(t, new(PersistenceTestSuite))
}
