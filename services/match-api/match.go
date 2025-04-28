package matchapi

// how are we going top represent high precision amouunts?
// min stake == 5 dollars
// initially only usdc supported
// chain? eth? sol? ...
// perhaps sol.

type MatchStatus string

// Constants
const (
	MatchOpen      MatchStatus = "open"
	MatchAccepted  MatchStatus = "accepted"
	MatchCompleted MatchStatus = "completed"
	MatchCancelled MatchStatus = "cancelled"
	MatchDisputed  MatchStatus = "disputed"
	MatchClaimed   MatchStatus = "claimed"
)

// vote is an address of the winner
type Player struct {
	walletAddress string
	vote          string
}

type match struct {
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
