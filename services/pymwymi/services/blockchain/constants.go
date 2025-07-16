package blockchain

const (
	WagerEscrowAddressBaseMainnet = "0xa2524aFF9536Ca744E93b25126D4Dced6198e948"
	WagerEscrowAddressBaseTestnet = "0xa2524aFF9536Ca744E93b25126D4Dced6198e948"
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
