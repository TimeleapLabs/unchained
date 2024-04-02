package pos

import (
	"math/big"
	"os"

	"github.com/KenshiTech/unchained/src/address"
	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/ethereum"
	"github.com/KenshiTech/unchained/src/ethereum/contracts"
	"github.com/KenshiTech/unchained/src/log"

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
	_, addrHex := address.CalculateHex(pkBytes[:])
	return GetVotingPower(addrHex, block)
}

func VotingPowerToFloat(power *big.Int) *big.Float {
	decimalPlaces := big.NewInt(1e18)
	powerFloat := new(big.Float).SetInt(power)
	powerFloat.Quo(powerFloat, new(big.Float).SetInt(decimalPlaces))
	return powerFloat
}

func Start() {
	base = big.NewInt(config.App.ProofOfStack.Base)

	pkBytes := bls.ClientPublicKey.Bytes()
	addrHexStr, addrHex := address.CalculateHex(pkBytes[:])

	log.Logger.
		With("Hex", addrHexStr).
		Info("Unchained")

	var err error

	posContract, err = ethereum.GetNewStakingContract(
		config.App.ProofOfStack.Chain,
		config.App.ProofOfStack.Address,
		false,
	)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to connect to the staking contract")

		os.Exit(1)
	}

	power, err := GetVotingPower(addrHex, big.NewInt(0))

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get voting power")

		return
	}

	total, err := GetTotalVotingPower()

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get total voting power")

		return
	}

	log.Logger.
		With("Power", VotingPowerToFloat(power)).
		With("Network", VotingPowerToFloat(total)).
		Info("PoS")
}

func init() {
	votingPowers = xsync.NewMapOf[[20]byte, *big.Int]()
	lastUpdated = xsync.NewMapOf[[20]byte, *big.Int]()
}
