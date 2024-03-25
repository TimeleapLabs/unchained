package pos

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/config"
	clientidentity "github.com/KenshiTech/unchained/crypto/client_identity"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/xerrors"

	"github.com/puzpuzpuz/xsync/v3"
)

var posContract *contracts.UnchainedStaking
var votingPowers *xsync.MapOf[[20]byte, *big.Int]
var lastUpdated *xsync.MapOf[[20]byte, *big.Int]
var base *big.Int

func GetTotalVotingPower() (*big.Int, error) {
	return posContract.GetTotalVotingPower(nil)
}

func GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error) {
	votingPower, err := posContract.GetVotingPower(nil, address)

	if err == nil {
		votingPowers.Store(address, votingPower)
		lastUpdated.Store(address, block)
	}

	return votingPower, err
}

func minBase(power *big.Int) *big.Int {
	if power == nil || power.Cmp(base) < 0 {
		return base
	}

	return power
}

func GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {
	powerLastUpdated, ok := lastUpdated.Load(address)

	if !ok {
		powerLastUpdated = big.NewInt(0)
	}

	updateAt := new(big.Int).Add(powerLastUpdated, big.NewInt(25000))

	if block.Cmp(updateAt) > 0 {
		votingPower, err := GetVotingPowerFromContract(address, block)
		return minBase(votingPower), err
	}

	if votingPower, ok := votingPowers.Load(address); ok {
		return minBase(votingPower), nil
	}

	return base, nil
}

func GetVotingPowerOfPublicKey(pkBytes [96]byte, block *big.Int) (*big.Int, error) {
	addr, err := address.NewAddress(pkBytes[:])
	if err != nil {
		return nil, err
	}
	return GetVotingPower(addr.Raw(), block)
}

func VotingPowerToFloat(power *big.Int) *big.Float {
	decimalPlaces := big.NewInt(1e18)
	powerFloat := new(big.Float).SetInt(power)
	powerFloat.Quo(powerFloat, new(big.Float).SetInt(decimalPlaces))
	return powerFloat
}

func Start() error {
	base = big.NewInt(config.Config.GetInt64("pos.base"))

	// todo remove the singleton pattern latter.
	pkBytes := clientidentity.GetPublicKey().Bytes()
	addr, err := address.NewAddress(pkBytes[:])
	if err != nil {
		return err
	}

	log.Logger.
		With("Hex", addr.Hex()).
		Info("Unchained")

	posContract, err = ethereum.GetNewStakingContract(
		config.Config.GetString("pos.chain"),
		config.Config.GetString("pos.address"),
		false,
	)

	if err != nil {
		return errors.Join(err, xerrors.ErrConnectionFailed("Failed to connect to the staking contract"))
	}

	power, err := GetVotingPower(addr.Raw(), big.NewInt(0))

	if err != nil {
		// todo log.Error
		log.Logger.
			With("Error", err).
			Error("Failed to get voting power")
		return errors.Join(err, fmt.Errorf("Failed to get voting power"))
	}

	total, err := GetTotalVotingPower()

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get total voting power")
		return err
	}

	log.Logger.
		With("Power", VotingPowerToFloat(power)).
		With("Network", VotingPowerToFloat(total)).
		Info("PoS")

	return nil
}

func init() {
	votingPowers = xsync.NewMapOf[[20]byte, *big.Int]()
	lastUpdated = xsync.NewMapOf[[20]byte, *big.Int]()
}
