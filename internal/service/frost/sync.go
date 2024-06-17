package frost

import (
	"context"
	"encoding/json"
	"math"
	"time"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/multisig"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// SendOnlineSigners sends online signers list to the Workers.
func (s *service) SendOnlineSigners(_ context.Context) error {
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
func (s *service) SyncSigners(ctx context.Context, onlineSigners []string) (<-chan *protocol.Message, error) {
	addresses, err := s.pos.GetSchnorrSigners(ctx)
	if err != nil {
		utils.Logger.With("Error", err).Error("Cant get signers list")
		return nil, err
	}

	addresses = FilterOnlineSigners(addresses, onlineSigners)

	minSignerCount := math.Round(0.65 * float64(len(addresses)))
	s.currentSigners = addresses
	var handshakeMessageCh <-chan *protocol.Message

	s.frost, handshakeMessageCh = multisig.NewIdentity(
		config.App.Plugins.Frost.Session,
		crypto.Identity.ExportEvmSigner().EvmAddress,
		addressArrayToStringArray(addresses),
		int(minSignerCount),
	)

	return handshakeMessageCh, nil
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
