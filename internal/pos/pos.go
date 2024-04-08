package pos

import (
	"math/big"
	"os"

	"github.com/KenshiTech/unchained/ethereum"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/KenshiTech/unchained/log"

	"github.com/puzpuzpuz/xsync/v3"
)

type Repository struct {
	ethRPC       *ethereum.Repository
	posContract  *contracts.UnchainedStaking
	votingPowers *xsync.MapOf[[20]byte, *big.Int]
	lastUpdated  *xsync.MapOf[[20]byte, *big.Int]
	base         *big.Int
}

func (s *Repository) GetTotalVotingPower() (*big.Int, error) {
	return s.posContract.GetTotalVotingPower(nil)
}

func (s *Repository) GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error) {
	votingPower, err := s.posContract.GetVotingPower(nil, address)

	if err == nil {
		s.votingPowers.Store(address, votingPower)
		s.lastUpdated.Store(address, block)
	}

	return votingPower, err
}

func (s *Repository) minBase(power *big.Int) *big.Int {
	if power == nil || power.Cmp(s.base) < 0 {
		return s.base
	}

	return power
}

func (s *Repository) GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {
	powerLastUpdated, ok := s.lastUpdated.Load(address)

	if !ok {
		powerLastUpdated = big.NewInt(0)
	}

	updateAt := new(big.Int).Add(powerLastUpdated, big.NewInt(25000))

	if block.Cmp(updateAt) > 0 {
		votingPower, err := s.GetVotingPowerFromContract(address, block)
		return s.minBase(votingPower), err
	}

	if votingPower, ok := s.votingPowers.Load(address); ok {
		return s.minBase(votingPower), nil
	}

	return s.base, nil
}

func (s *Repository) GetVotingPowerOfPublicKey(pkBytes [96]byte, block *big.Int) (*big.Int, error) {
	_, addrHex := address.CalculateHex(pkBytes[:])
	return s.GetVotingPower(addrHex, block)
}

func (s *Repository) VotingPowerToFloat(power *big.Int) *big.Float {
	decimalPlaces := big.NewInt(1e18)
	powerFloat := new(big.Float).SetInt(power)
	powerFloat.Quo(powerFloat, new(big.Float).SetInt(decimalPlaces))
	return powerFloat
}

func New(
	ethRPC *ethereum.Repository,
) *Repository {
	s := &Repository{ethRPC: ethRPC}
	s.init()

	s.base = big.NewInt(config.App.ProofOfStake.Base)

	pkBytes := bls.ClientPublicKey.Bytes()
	addrHexStr, addrHex := address.CalculateHex(pkBytes[:])

	log.Logger.
		With("Hex", addrHexStr).
		Info("Unchained")

	var err error

	s.posContract, err = s.ethRPC.GetNewStakingContract(
		config.App.ProofOfStake.Chain,
		config.App.ProofOfStake.Address,
		false,
	)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to connect to the staking contract")

		os.Exit(1)
	}

	power, err := s.GetVotingPower(addrHex, big.NewInt(0))

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get voting power")

		return s
	}

	total, err := s.GetTotalVotingPower()

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get total voting power")

		return s
	}

	log.Logger.
		With("Power", s.VotingPowerToFloat(power)).
		With("Network", s.VotingPowerToFloat(total)).
		Info("PoS")

	return s
}

func (s *Repository) init() {
	s.votingPowers = xsync.NewMapOf[[20]byte, *big.Int]()
	s.lastUpdated = xsync.NewMapOf[[20]byte, *big.Int]()
}
