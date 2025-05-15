package pymwymi

// how are we going top represent high precision amouunts?
// min stake == 5 dollars
// initially only usdc supported
// chain? eth? sol? ...
// perhaps sol.

type ChallengeStatus int8

// Constants
const (
	StateCreated ChallengeStatus = iota
	StatePending
	StateCancelled
	StateCompleted
	StateClaimed
)

type WalletAddress string

// vote is an address of the winner
type Player struct {
	name          string
	walletAddress WalletAddress
	vote          string
}

type Challenge struct {
	creator      WalletAddress
	name         string
	category     string
	description  string
	location     string
	stake        int
	currency     string
	participants []Player
	status       ChallengeStatus
	winner       string
}
