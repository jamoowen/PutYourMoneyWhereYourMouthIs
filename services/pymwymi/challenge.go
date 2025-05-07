package pymwymi

// how are we going top represent high precision amouunts?
// min stake == 5 dollars
// initially only usdc supported
// chain? eth? sol? ...
// perhaps sol.

type ChallengeStatus string

// Constants
const (
	ChallengeOpen      ChallengeStatus = "open"
	ChallengeAccepted  ChallengeStatus = "accepted"
	ChallengeCompleted ChallengeStatus = "completed"
	ChallengeCancelled ChallengeStatus = "cancelled"
	ChallengeDisputed  ChallengeStatus = "disputed"
	ChallengeClaimed   ChallengeStatus = "claimed"
)

// vote is an address of the winner
type Player struct {
	walletAddress string
	vote          string
}

type Challenge struct {
	creator     string
	challenger  string
	name        string
	category    string
	description string
	location    string
	stake       int
	currency    string
	status      string
	winner      string
}
