package frost

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/multisig"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// SyncSigners starts calculating of Frost signers by sending signers list to the Broker.
func (s *service) SyncSigners(ctx context.Context) error {
	addresses, err := s.pos.GetSchnorrSigners(ctx)
	if err != nil {
		utils.Logger.With("Error", err).Error("Cant get signers list")
		return err
	}

	minSignerCount := (len(addresses) / 2) + 1
	s.currentSigners = addresses

	var handshakeMessageCh <-chan *protocol.Message
	crypto.Identity.Frost, handshakeMessageCh = multisig.NewIdentity(
		config.App.Plugins.Frost.Session,
		crypto.Identity.ExportEvmSigner().EvmAddress,
		addressArrayToStringArray(addresses),
		minSignerCount,
	)

	go s.CoordinateHandshake(handshakeMessageCh)

	return nil
}

func (s *service) CoordinateHandshake(ch <-chan *protocol.Message) {
	for msg := range ch {
		msgBytes, err := msg.MarshalBinary()
		if err != nil {
			utils.Logger.With("Error", err).Error("Cant marshal handshake message")
		}

		conn.Send(consts.OpCodeFrostSignerHandshake, msgBytes)
	}
}

func addressArrayToStringArray(addresses []common.Address) []string {
	strAddresses := make([]string, 0, len(addresses))
	for _, address := range addresses {
		strAddresses = append(strAddresses, address.String())
	}

	return strAddresses
}
