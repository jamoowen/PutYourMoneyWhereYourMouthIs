package blockchain

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

type BlockchainService interface {
	AuthenticateSignature(walletAddress string, signatureHex string) error
}

type Service struct{}

func NewBlockchainService() *Service {
	return &Service{}
}

func GetWagerEscrowAddress(env ChainEnvironment, chain ChainName) string {
	switch chain {
	case BaseChain:
		if env == Mainnet {
			return WagerEscrowAddressBaseMainnet
		}
		return WagerEscrowAddressBaseTestnet
	default:
		return ""
	}
}

func GetChainID(env ChainEnvironment) int64 {
	if env == Mainnet {
		return BaseMainnetChainID
	}
	return BaseTestnetChainID
}

func AuthenticateSignature(walletAddress string, signatureHex string, originalMessage string) (bool, *pymwymi.Error) {
	signatureHex = strings.TrimPrefix(signatureHex, "0x")
	sigBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, pymwymi.Errorf(pymwymi.ErrBadInput, "failed to decode signature: %s", err.Error())
	}
	if len(sigBytes) != 65 {
		return false, pymwymi.Errorf(pymwymi.ErrBadInput, "signature must be 65 bytes long, got %d", len(sigBytes))
	}

	// ⚠️ Normalize v from [27,28] → [0,1]
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	msgHash := accounts.TextHash([]byte(originalMessage))

	pubKeyFromSig, err := crypto.SigToPub(msgHash, sigBytes)
	if err != nil {
		return false, pymwymi.Errorf(pymwymi.ErrBadInput, "failed to recreate public key from signature: %s", err.Error())
	}

	walletAddressFromPubKey := crypto.PubkeyToAddress(*pubKeyFromSig).Hex()
	if !strings.EqualFold(walletAddress, walletAddressFromPubKey) {
		return false, nil
	}

	return true, nil
}

func (s *Service) ToWeiUSDC(value int64) int64 {
	return value * 1_000_000 // 10^6
}

func (s *Service) FromWeiUSDC(value int64) int64 {
	return value / 1_000_000
}
