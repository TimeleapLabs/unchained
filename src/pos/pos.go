package pos

import (
	"math/big"
	"os"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/KenshiTech/unchained/log"
)

var posContract *contracts.UnchainedStaking
var votingPowers map[[20]byte]*big.Int
var stakes map[[20]byte]contracts.UnchainedStakingStake

func GetTotalVotingPower() (*big.Int, error) {
	return posContract.GetTotalVotingPower(nil)
}

func GetStake(address [20]byte, block *big.Int) (*big.Int, error) {
	cachedStake, ok := stakes[address]

	if ok && cachedStake.Unlock.Cmp(block) >= 0 {
		cachedPower, ok := votingPowers[address]

		if ok {
			return cachedPower, nil
		}
	}

	stake, err := posContract.GetStake0(nil, address)

	if err == nil {
		if stake.Amount.Cmp(big.NewInt(0)) == 0 {
			// TODO: we should listen to stake changed events
			// TODO: and update stakes accordingly
			const numOfBlocks int64 = 25000
			stake.Unlock = new(big.Int).Add(block, big.NewInt(numOfBlocks))
		}

		stakes[address] = stake
	}

	votingPower, err := posContract.GetVotingPower0(nil, address)

	if err != nil {
		votingPowers[address] = votingPower
	}

	return votingPower, err
}

func maxBase(power *big.Int) *big.Int {
	base := config.Config.GetUint64("pos.base")
	baseBig := big.NewInt(int64(base))

	if power.Cmp(baseBig) < 0 {
		return baseBig
	}

	return power
}

func GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {
	if votingPower, ok := votingPowers[address]; ok {
		return maxBase(votingPower), nil
	}

	votingPower, err := GetStake(address, block)

	if err != nil {
		return nil, err
	}

	return maxBase(votingPower), nil
}

func GetVotingPowerOfPublicKey(
	pkBytes [96]byte,
	block *big.Int,
) (*big.Int, error) {
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
	pkBytes := bls.ClientPublicKey.Bytes()
	addrHexStr, addrHex := address.CalculateHex(pkBytes[:])

	log.Logger.
		With("Hex", addrHexStr).
		Info("Unchained")

	var err error

	posContract, err = ethereum.GetNewStakingContract(
		config.Config.GetString("pos.chain"),
		config.Config.GetString("pos.address"),
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
	votingPowers = make(map[[20]byte]*big.Int)
	stakes = make(map[[20]byte]contracts.UnchainedStakingStake)
}
