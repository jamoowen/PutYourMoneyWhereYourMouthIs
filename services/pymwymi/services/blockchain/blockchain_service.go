package blockchain

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type BlockchainService interface {
	AuthenticateSignature(walletAddress string, signatureHex string) error
}

type Service struct{}

func NewBlockchainService() *Service {
	return &Service{}
}

func AuthenticateSignature(walletAddress string, signatureHex string, originalMessage string) (bool, error) {
	// 1. get signature bytes
	sigBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}
	// 2. recreate pub key from sig
	pubKeyFromSig, err := crypto.SigToPub(crypto.Keccak256Hash([]byte(originalMessage)).Bytes(), sigBytes)
	if err != nil {
		return false, fmt.Errorf("failed to recreate public key from signature: %w", err)
	}
	// 3. check if wallet sent matches the extracted wallet
	walletAddressFromPubKey := crypto.PubkeyToAddress(*pubKeyFromSig).Hex()
	if !strings.EqualFold(walletAddress, walletAddressFromPubKey) {
		return false, nil
	}
	return true, nil
}
