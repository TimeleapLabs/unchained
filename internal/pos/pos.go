package pos

import (
	"math/big"
	"os"

	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/TimeleapLabs/unchained/internal/address"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/pos/eip712"

	"github.com/puzpuzpuz/xsync/v3"
)

type Repository struct {
	ethRPC       *ethereum.Repository
	posContract  *contracts.UnchainedStaking
	posChain     string
	votingPowers *xsync.MapOf[[20]byte, *big.Int]
	lastUpdated  *xsync.MapOf[[20]byte, *big.Int]
	base         *big.Int
	eip712Signer *eip712.Signer
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

func (s *Repository) GetVotingPowerOfPublicKeyAtBlock(pkBytes [96]byte, block *big.Int) (*big.Int, error) {
	_, addrHex := address.CalculateHex(pkBytes[:])
	return s.GetVotingPower(addrHex, block)
}

func (s *Repository) GetVotingPowerOfPublicKey(pkBytes [96]byte) (*big.Int, error) {
	_, addrHex := address.CalculateHex(pkBytes[:])
	block, err := s.ethRPC.GetBlockNumber(s.posChain)
	if err != nil {
		return nil, err
	}
	return s.GetVotingPower(addrHex, big.NewInt(int64(block)))
}

func (s *Repository) VotingPowerToFloat(power *big.Int) *big.Float {
	decimalPlaces := big.NewInt(1e18)
	powerFloat := new(big.Float).SetInt(power)
	powerFloat.Quo(powerFloat, new(big.Float).SetInt(decimalPlaces))
	return powerFloat
}

func New(ethRPC *ethereum.Repository) *Repository {
	s := &Repository{
		ethRPC: ethRPC,
	}

	s.init()

	s.base = big.NewInt(config.App.ProofOfStake.Base)

	pkBytes := crypto.Identity.Bls.PublicKey.Bytes()
	addrHexStr, addrHex := address.CalculateHex(pkBytes[:])

	log.Logger.
		With("Address", addrHexStr).
		Info("PoS identity initialized")

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

	chainID, err := s.posContract.GetChainId(nil)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get chain ID")

		return s
	}

	s.eip712Signer = eip712.New(chainID, config.App.ProofOfStake.Address)
	s.posChain = config.App.ProofOfStake.Chain

	return s
}

func (s *Repository) init() {
	s.votingPowers = xsync.NewMapOf[[20]byte, *big.Int]()
	s.lastUpdated = xsync.NewMapOf[[20]byte, *big.Int]()
}
