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

type Player struct {
	WalletAddress       WalletAddress `bson:"walletAddress" json:"walletAddress"`
	Vote                string        `bson:"vote" json:"vote"`
	VoteTransactionHash string        `bson:"voteTransactionHash" json:"voteTransactionHash"`
}

type NewChallengeDto struct {
	TransactionHash       string   `json:"transactionHash"`
	Creator               string   `json:"creator"`
	Name                  string   `json:"name"`
	Category              string   `json:"category"`
	Description           string   `json:"description"`
	Location              string   `json:"location"`
	Stake                 string   `json:"stake"`
	Currency              string   `json:"currency"`
	ParticipantsAddresses []string `json:"participantsAddresses"`
}

type Challenge struct {
	TransactionHash string          `bson:"transactionHash" json:"transactionHash"`
	Creator         WalletAddress   `bson:"creator" json:"creator"`
	Name            string          `bson:"name" json:"name"`
	Category        string          `bson:"category" json:"category"`
	Description     string          `bson:"description" json:"description"`
	Location        string          `bson:"location" json:"location"`
	Stake           string          `bson:"stake" json:"stake"`
	Currency        string          `bson:"currency" json:"currency"`
	Participants    []Player        `bson:"participants" json:"participants"`
	Status          ChallengeStatus `bson:"status" json:"status"`
	Winner          string          `bson:"winner" json:"winner"`
}

type PersistedChallenge struct {
	ID        string `json:"_id" bson:"_id"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	Challenge
}
