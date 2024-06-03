package pos

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"
	"github.com/TimeleapLabs/unchained/internal/service/pos/eip712"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/TimeleapLabs/unchained/internal/utils/address"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
)

type Service interface {
	GetTotalVotingPower() (*big.Int, error)
	GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error)
	GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error)
	GetVotingPowerOfEvm(ctx context.Context, evmAddress string) (*big.Int, error)
	GetVotingPowerOfPublicKey(ctx context.Context, pkBytes [96]byte) (*big.Int, error)
	GetSchnorrSigners(ctx context.Context) ([]common.Address, error)
}

type service struct {
	ethRPC       ethereum.RPC
	posContract  *contracts.ProofOfStake
	votingPowers *xsync.MapOf[[20]byte, *big.Int]
	lastUpdated  *xsync.MapOf[[20]byte, *big.Int]
	base         *big.Int
	eip712Signer *eip712.Signer
}

func (s *service) GetTotalVotingPower() (*big.Int, error) {
	return new(big.Int).Mul(big.NewInt(5e10), big.NewInt(1e18)), nil
	// return s.posContract.GetTotalVotingPower(nil)
}

func (s *service) GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error) {
	stake, err := s.posContract.GetStake(nil, address)
	// votingPower, err := s.posContract.GetVotingPower(nil, address)
	if err != nil {
		return stake.Amount, err
	}

	s.votingPowers.Store(address, stake.Amount)
	s.lastUpdated.Store(address, block)

	return stake.Amount, nil
}

func (s *service) minBase(power *big.Int) *big.Int {
	if power == nil || power.Cmp(s.base) < 0 {
		return s.base
	}

	return power
}

func (s *service) GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {
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

func (s *service) GetVotingPowerOfEvm(ctx context.Context, evmAddress string) (*big.Int, error) {
	block, err := s.ethRPC.GetBlockNumber(ctx, config.App.ProofOfStake.Chain)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(evmAddress)
	return s.GetVotingPower(address, big.NewInt(int64(block)))
}

func (s *service) GetSchnorrSigners(ctx context.Context) ([]common.Address, error) {
	return s.posContract.GetValidators(&bind.CallOpts{Context: ctx})
}

func (s *service) GetVotingPowerOfPublicKey(ctx context.Context, pkBytes [96]byte) (*big.Int, error) {
	_, addrHex := address.CalculateHex(pkBytes[:])
	block, err := s.ethRPC.GetBlockNumber(ctx, config.App.ProofOfStake.Chain)
	if err != nil {
		return nil, err
	}
	return s.GetVotingPower(addrHex, big.NewInt(int64(block)))
}

func New(ethRPC ethereum.RPC) Service {
	s := &service{
		ethRPC:       ethRPC,
		base:         big.NewInt(config.App.ProofOfStake.Base),
		votingPowers: xsync.NewMapOf[[20]byte, *big.Int](),
		lastUpdated:  xsync.NewMapOf[[20]byte, *big.Int](),
	}

	blsPublicKey, err := hex.DecodeString(config.App.Secret.PublicKey)
	if err != nil {
		panic(err)
	}

	s.posContract, err = s.ethRPC.GetNewStakingContract(
		config.App.ProofOfStake.Chain,
		config.App.ProofOfStake.Address,
		false,
	)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to connect to the staking contract")

		panic(err)
	}

	power, err := s.GetVotingPower([20]byte(blsPublicKey), big.NewInt(0))
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to get voting power")

		panic(err)
	}

	total, err := s.GetTotalVotingPower()
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to get total voting power")

		panic(err)
	}

	utils.Logger.
		With("Power", utils.BigIntToFloat(power)).
		With("Network", utils.BigIntToFloat(total)).
		Info("PoS")

	// chainID, err := s.posContract.GetChainId(nil)
	// if err != nil {
	// 	utils.Logger.
	// 		With("Error", err).
	// 		Error("Failed to get chain ID")

	// 	panic(err)
	// }

	chainID := big.NewInt(421614)
	s.eip712Signer = eip712.New(chainID, config.App.ProofOfStake.Address)

	return s
}
