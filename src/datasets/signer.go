package datasets

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

func (m *Signer) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}
