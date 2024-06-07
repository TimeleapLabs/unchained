package multisig

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/frost"
)

// DistributedSigner represents a MultiSig identity.
type DistributedSigner struct {
	minSigningCount int

	ID         party.ID
	sessionID  string
	Signers    []party.ID
	ackHandler *protocol.MultiHandler
	Config     *frost.TaprootConfig
}

// ConfirmFromBytes function will set other parties confirms from byte array.
func (d *DistributedSigner) ConfirmFromBytes(msgBytes []byte) (bool, error) {
	msg := &protocol.Message{}
	err := msg.UnmarshalBinary(msgBytes)
	if err != nil {
		utils.Logger.With("err", err).Error("cant unmarshal message")
		return false, consts.ErrCantDecode
	}

	return d.Confirm(msg)
}

// Confirm function will set other parties confirms.
func (d *DistributedSigner) Confirm(msg *protocol.Message) (bool, error) {
	if !msg.IsFor(d.ID) {
		return false, consts.ErrInvalidSignature
	}

	d.ackHandler.Accept(msg)

	result, err := d.ackHandler.Result()
	if err != nil {
		if err.Error() == "protocol: not finished" {
			return false, nil
		}

		return false, err
	}

	var ok bool
	d.Config, ok = result.(*frost.TaprootConfig)
	if !ok {
		utils.Logger.Error("Can't get TaprootConfig")
		return false, consts.ErrInternalError
	}

	return true, nil
}

// NewIdentity creates a new MultiSig identity.
func NewIdentity(sessionID, id string, signers []string, minSigningCount int) (*DistributedSigner, <-chan *protocol.Message) {
	signersIDs := NewSignersFromStrings(signers)

	startSession := frost.KeygenTaproot(party.ID(id), signersIDs, minSigningCount)
	handler, err := protocol.NewMultiHandler(startSession, []byte(sessionID))
	if err != nil {
		panic("Cant create multi-sig identity: " + err.Error())
	}

	return &DistributedSigner{
		sessionID:       sessionID,
		Signers:         signersIDs,
		minSigningCount: minSigningCount,
		ackHandler:      handler,
	}, handler.Listen()
}
