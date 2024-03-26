package datasets

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

func NewChallengeWithRandom(payload [128]byte) Challenge {
	return Challenge{
		Random: payload[:],
	}
}

func NewChallenge(payload []byte) (*Challenge, error) {
	var challenge Challenge

	err := proto.Unmarshal(payload, &challenge)
	if err != nil {
		return nil, constants.ErrInternalError
	}

	return &challenge, nil
}

func (m *Challenge) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}
