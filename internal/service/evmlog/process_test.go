package evmlog

import (
	"context"
)

func (s *EvmLogTestSuite) TestProcessBlocks() {
	err := s.service.ProcessBlocks(context.Background(), "eth")
	s.NoError(err)
}
