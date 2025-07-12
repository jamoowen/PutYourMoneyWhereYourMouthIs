package blockchain

const (
	WagerEscrowAddressBaseMainnet = "0xbc7CA0b61Eb1B0c22F7D93CEC3C993D6d7079645"
	WagerEscrowAddressBaseTestnet = "0xbc7CA0b61Eb1B0c22F7D93CEC3C993D6d7079645"
)

type ChainName string

const (
	BaseChain ChainName = "base"
)

type ChainEnvironment string

const (
	Mainnet ChainEnvironment = "mainnet"
	Testnet ChainEnvironment = "testnet"
)

const (
	BaseMainnetChainID int64 = 8453
	BaseTestnetChainID int64 = 84532
)
