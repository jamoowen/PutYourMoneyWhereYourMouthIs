package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/contracts"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" {
		printHelp()
		return
	}

	// Parse CLI args
	var env blockchain.ChainEnvironment = blockchain.Testnet
	envFlag := args[0]
	if envFlag != "main" && envFlag != "test" {
		log.Fatal("First argument must be 'main' or 'test'")
	}
	if envFlag == "main" {
		env = blockchain.Mainnet
	}

	if len(args) < 2 {
		log.Fatal("Missing command")
	}
	command := args[1]
	params := args[2:]

	// Contract + RPC Setup
	rpcURL, privateKeyHex := validateEnvVars(env)

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	defer client.Close()

	chainID := blockchain.GetChainID(env)
	contractAddress := blockchain.GetWagerEscrowAddress(env, blockchain.BaseChain)
	address := common.HexToAddress(contractAddress)
	instance, err := contracts.NewWagerEscrow(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Command dispatch
	switch command {
	case "tokens":
		handleGetSupportedTokens(instance)
	case "add-token":
		if len(params) < 1 {
			log.Fatal("Missing token address")
		}
		auth := prepareForContractWrite(client, chainID, privateKeyHex)
		handleAddSupportedToken(auth, instance, params[0])
	case "wager":
		if len(params) < 1 {
			log.Fatal("Missing wager ID")
		}
		ID, err := strconv.Atoi(params[0])
		if err != nil {
			log.Fatal(err)
		}
		handleGetWager(instance, ID)
	case "all-wagers":
		handleGetAllWagerIDs(instance)
	default:
		fmt.Println("Unknown command:", command)
		printHelp()
	}
}

func handleGetWager(instance *contracts.WagerEscrow, ID int) {
	bigID := big.NewInt(int64(ID))
	result, err := instance.GetWager(nil, bigID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wager:", result)
}

func handleGetAllWagerIDs(instance *contracts.WagerEscrow) {
	result, err := instance.GetAllWagers(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wager ids:", result)
}

func handleGetSupportedTokens(instance *contracts.WagerEscrow) {
	supportedTokens, err := instance.GetSupportedTokens(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Supported tokens:", supportedTokens)
}

func handleAddSupportedToken(auth *bind.TransactOpts, instance *contracts.WagerEscrow, tokenAddressHex string) {
	fmt.Printf("Adding %s to supported tokens\n", tokenAddressHex)
	tokenAddress := common.HexToAddress(tokenAddressHex)
	tx, err := instance.AddToken(auth, tokenAddress)
	if err != nil {
		log.Fatalf("Failed to add token: %v", err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func prepareForContractWrite(client *ethclient.Client, chainID int64, privateKeyHex string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatalf("invalid PRIVATE_KEY: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to retrieve nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("failed to suggest gas price: %v", err)
	}
	// auth := bind.NewKeyedTransactor(privateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("failed to create auth: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth
}

func validateEnvVars(env blockchain.ChainEnvironment) (string, string) {
	rpcURL := os.Getenv("BASE_TESTNET_URL")
	if env == blockchain.Mainnet {
		rpcURL = os.Getenv("BASE_MAINNET_URL")
	}
	if rpcURL == "" {
		log.Fatal("missing BASE_TESTNET_URL or BASE_MAINNET_URL")
	}
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("missing PRIVATE_KEY")
	}
	return rpcURL, privateKey
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./cmd/cli/ [env] [command] [params...]")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run ./cmd/cli/ test tokens")
	fmt.Println("  go run ./cmd/cli/ main add-token 0xTOKENADDRESS")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  tokens             - List supported tokens")
	fmt.Println("  add-token [token] - Add a token to supported list")
	fmt.Println("  wager [id]        - Get a wager by ID")
	fmt.Println("  all-wagers        - Get all wager IDs")
}
