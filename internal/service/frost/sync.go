package frost

import (
	"github.com/TimeleapLabs/unchained/internal/model"
)

// SyncFrost starts calculating of Frost signers by sending signers list to the Broker.
func (s *service) SyncFrost() error {
	// signers := []model.Signer{}
	// store.Signers.Range(func(_ *websocket.Conn, value model.Signer) bool {
	//	signers = append(signers, value)
	//	return true
	// })
	//
	// signersBytes, err := json.Marshal(signers)
	//if err != nil {
	//	utils.Logger.With("Error", err).Error("Cant marshal signers list")
	//	return consts.ErrInternalError
	//}

	return nil
}

// SyncSigners Get list of signers and check power of voting them and generate a new list (if there is difference) of signers which have power.
func (s *service) SyncSigners(signers []model.Signer) error {
	// TODO: get power of list items and delete no power ones.

	// TODO: check the final list with previous one, and replace it if it have difference

	return nil
}
