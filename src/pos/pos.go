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

func GetTotalVotingPower() (*big.Int, error) {
	return posContract.TotalVotingPower(nil)
}

func GetVotingPower(address [20]byte) (contracts.UnchainedStakingStake, error) {
	stake, err := posContract.StakeOf0(nil, address)
	return stake, err
}

func GetVotingPowerOfPublicKey(pkBytes [96]byte) (contracts.UnchainedStakingStake, error) {
	_, addrHex := address.CalculateHex(pkBytes[:])
	stake, err := posContract.StakeOf0(nil, addrHex)
	return stake, err
}

func Start() {

	pkBytes := bls.ClientPublicKey.Bytes()
	addrHexStr, addrHex := address.CalculateHex(pkBytes[:])

	log.Logger.
		With("Link", addrHexStr).
		Info("Unchained EVM")

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

	stake, _ := GetVotingPower(addrHex)
	total, _ := GetTotalVotingPower()

	var votingPower *big.Int

	base := big.NewInt(0)
	base.SetString(config.Config.GetString("pos.base"), 10)

	nft := big.NewInt(0)
	nft.SetString(config.Config.GetString("pos.nft"), 10)

	if err != nil {
		votingPower = base
	} else {
		tokenPower := new(big.Int).Add(stake.Amount, base)
		nftPower := new(big.Int).Mul(nft, big.NewInt(int64(len(stake.NftIds))))
		votingPower = new(big.Int).Add(tokenPower, nftPower)
	}

	log.Logger.
		With("Power", votingPower.String()).
		With("Network", total.String()).
		Info("Voting power")
}
