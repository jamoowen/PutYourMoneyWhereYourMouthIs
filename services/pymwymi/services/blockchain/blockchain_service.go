package blockchain

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

type BlockchainService interface {
	AuthenticateSignature(walletAddress string, signatureHex string) error
}

type Service struct{}

func AuthenticateSignature(walletAddress string, signatureHex string) (bool, error) {
	// 1. get signature bytes
	sigBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, err
	}
	// 2. recreate pub key from sig
	pubKeyFromSig, err := crypto.SigToPub(crypto.Keccak256Hash([]byte(pymwymi.SIGN_IN_STRING)).Bytes(), sigBytes)
	if err != nil {
		return false, err
	}
	// 3. check if wallet sent matches the extracted wallet
	walletAddressFromPubKey := crypto.PubkeyToAddress(*pubKeyFromSig).Hex()
	if !strings.EqualFold(walletAddress, walletAddressFromPubKey) {
		return false, nil
	}
	return true, nil
}
