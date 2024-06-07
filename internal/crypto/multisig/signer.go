package multisig

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/frost"
)

// MessageSigner represents state of signing of a message.
type MessageSigner struct {
	ackHandler *protocol.MultiHandler
}

// Confirm function will set other parties confirms.
func (s *MessageSigner) Confirm(msgByte []byte) error {
	msg := &protocol.Message{}
	err := msg.UnmarshalBinary(msgByte)
	if err != nil {
		utils.Logger.With("err", err).Error("cant unmarshal message")
		return consts.ErrCantDecode
	}

	s.ackHandler.Accept(msg)

	return nil
}

func (d *DistributedSigner) NewSigner(data []byte) (*MessageSigner, <-chan *protocol.Message, error) {
	if d.Config == nil {
		return nil, nil, consts.ErrSignerIsNotReady
	}

	dataKeccak256 := Keccak256(data)

	startSession := frost.SignTaproot(d.Config, d.Signers, dataKeccak256)
	handler, err := protocol.NewMultiHandler(startSession, []byte(d.sessionID))
	if err != nil {
		panic(err)
	}

	return &MessageSigner{
		ackHandler: handler,
	}, handler.Listen(), nil
}

// func decodeMessages(in <-chan *protocol.Message) <-chan []byte {
//	out := make(chan []byte)
//
//	go func() {
//		for msg := range in {
//			msgByte, err := msg.MarshalBinary()
//			if err != nil {
//				utils.Logger.With("err", err).Error("cant marshal message")
//				continue
//			}
//
//			out <- msgByte
//		}
//
//		close(out)
//	}()
//
//	return out
//}
