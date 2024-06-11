package frost

import (
	"context"
	"encoding/json"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/multisig"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

func (s *service) SendOnlineSigners(ctx context.Context) error {
	onlineSigners := []string{}
	store.OnlineFrostParties.Range(func(key string, value time.Time) bool {
		onlineSigners = append(onlineSigners, key)

		return false
	})

	onlineSignersBytes, err := json.Marshal(onlineSigners)
	if err != nil {
		utils.Logger.With("Error", err).Error("Cant marshal online signers list")
		return consts.ErrInternalError
	}

	pubsub.Publish(consts.ChannelFrostSigner, consts.OpcodeFrostSignerOnlines, onlineSignersBytes)

	return nil
}

// SyncSigners starts calculating of Frost signers by sending signers list to the Broker.
func (s *service) SyncSigners(ctx context.Context, onlineSigners []string) error {
	addresses, err := s.pos.GetSchnorrSigners(ctx)
	if err != nil {
		utils.Logger.With("Error", err).Error("Cant get signers list")
		return err
	}

	addresses = FilterOnlineSigners(addresses, onlineSigners)

	minSignerCount := (len(addresses) / 100) * 65
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
