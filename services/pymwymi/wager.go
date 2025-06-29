package pymwymi

// how are we going top represent high precision amouunts?
// min stake == 5 dollars
// initially only usdc supported
// chain? eth? sol? ...
// perhaps sol.

type WagerStatus int8

// Constants
const (
	StateCreated WagerStatus = iota
	StatePending
	StateCancelled
	StateCompleted
	StateClaimed
)

type VoteIntent string

const (
	VotePending VoteIntent = "pending"
	VoteCancel  VoteIntent = "cancel"
	VoteWinner  VoteIntent = "winner"
)

type Vote struct {
	HasVoted bool       `bson:"hasVoted" json:"hasVoted"`
	Intent   VoteIntent `bson:"intent" json:"intent"`
	Winner   string     `bson:"winner" json:"winner"`
}

type Player struct {
	WalletAddress string `bson:"walletAddress" json:"walletAddress"`
	Vote          Vote   `bson:"vote" json:"vote"`
	HasStaked     bool   `bson:"hasStaked" json:"hasStaked"`
	HasWithdrawn  bool   `bson:"hasWithdrawn" json:"hasWithdrawn"`
}

type Wager struct {
	ID              string      `bson:"_id" json:"_id"`
	TransactionHash string      `bson:"transactionHash" json:"transactionHash"`
	Creator         string      `bson:"creator" json:"creator"`
	Name            string      `bson:"name" json:"name"`
	Category        string      `bson:"category" json:"category"`
	Description     string      `bson:"description" json:"description"`
	Location        string      `bson:"location" json:"location"`
	Stake           string      `bson:"stake" json:"stake"`
	Currency        string      `bson:"currency" json:"currency"`
	Participants    []Player    `bson:"participants" json:"participants"`
	Status          WagerStatus `bson:"status" json:"status"`
	Winner          string      `bson:"winner" json:"winner"`
}

type PersistedWager struct {
	CreatedAt int64 `json:"createdAt" bson:"createdAt"`
	Wager
}
