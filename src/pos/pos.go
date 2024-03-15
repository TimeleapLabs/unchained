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
var votingPowers map[[20]byte]big.Int
var stakes map[[20]byte]contracts.UnchainedStakingStake

func GetTotalVotingPower() (*big.Int, error) {
	return posContract.TotalVotingPower(nil)
}

func GetStake(address [20]byte, block *big.Int) (contracts.UnchainedStakingStake, error) {
	cached, ok := stakes[address]

	if ok && cached.Unlock.Cmp(block) >= 0 {
		return cached, nil
	}

	stake, err := posContract.StakeOf0(nil, address)

	if err == nil {
		if stake.Amount.Cmp(big.NewInt(0)) == 0 {
			// TODO: we should listen to stake changed events
			// TODO: and update stakes accordingly
			stake.Unlock = new(big.Int).Add(block, big.NewInt(25000))
		}

		stakes[address] = stake
	}

	return stake, err
}

func GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {

	if votingPower, ok := votingPowers[address]; ok {
		return &votingPower, nil
	}

	stake, err := GetStake(address, block)

	if err != nil {
		return nil, err
	}

	base := big.NewInt(0)
	base.SetString(config.Config.GetString("pos.base"), 10)

	nft := big.NewInt(0)
	nft.SetString(config.Config.GetString("pos.nft"), 10)

	nftPower := new(big.Int).Mul(nft, big.NewInt(int64(len(stake.NftIds))))
	votingPower := new(big.Int).Add(stake.Amount, nftPower)

	if votingPower.Cmp(base) < 0 {
		votingPower = base
	}

	votingPowers[address] = *votingPower
	return votingPower, nil
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
	votingPowers = make(map[[20]byte]big.Int)
	stakes = make(map[[20]byte]contracts.UnchainedStakingStake)
}
