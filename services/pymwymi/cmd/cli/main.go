package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	env := args[0]
	if env != "dev" && env != "prod" {
		log.Fatal("First argument must be 'dev' or 'prod'")
	}

	if len(args) < 2 {
		log.Fatal("Missing command")
	}
	command := args[1]
	params := args[2:]

	// Contract + RPC Setup
	var contractAddress string
	if env == "dev" {
		contractAddress = "0x7c471fcf09959b8522760ca69bddf3c91900d834"
	} else if env == "prod" {
		contractAddress = "0x7c471fcf09959b8522760ca69bddf3c91900d834"
	}

	rpcURL := os.Getenv("BASE_TESTNET_URL")
	if env == "prod" {
		rpcURL = os.Getenv("BASE_MAINNET_URL")
	}
	if rpcURL == "" {
		log.Fatal("Missing BASE_TESTNET_URL or BASE_MAINNET_URL")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	defer client.Close()

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("Missing PRIVATE_KEY")
	}
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatalf("Invalid PRIVATE_KEY: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(84532)) // Replace with real chain ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Command dispatch
	switch command {
	case "supported-tokens":
		handleSupportedTokens(client, contractAddress)
	case "add-supported-token":
		if len(params) < 1 {
			log.Fatal("Missing token address")
		}
		handleAddSupportedToken(client, auth, contractAddress, params[0])
	default:
		fmt.Println("Unknown command:", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./cmd/cli/ [env] [command] [params...]")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run ./cmd/cli/ dev supported-tokens")
	fmt.Println("  go run ./cmd/cli/ dev add-supported-token 0xTOKENADDRESS")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  supported-tokens             - List supported tokens")
	fmt.Println("  add-supported-token [token] - Add a token to supported list")
}
