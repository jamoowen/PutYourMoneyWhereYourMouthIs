package pymwymi

import "go.mongodb.org/mongo-driver/v2/bson"

// how are we going top represent high precision amouunts?
// min stake == 5 dollars
// initially only usdc supported
// chain? eth? sol? ...
// perhaps sol.

type (
	WagerStatus       int
	InteractionStatus string
	VoteIntent        string
)

// these map directly to the states on the smart contract
const (
	StateCreated WagerStatus = iota
	StatePending
	StateCancelled
	StateCompleted
	StateClaimed
)

const (
	InteractionStateDormant   InteractionStatus = "dormant"
	InteractionStatePending   InteractionStatus = "pending"
	InteractionStateConfirmed InteractionStatus = "confirmed"
)

const (
	VoteUnknown VoteIntent = "unknown"
	VoteCancel  VoteIntent = "cancel"
	VoteWinner  VoteIntent = "winner"
)

type Vote struct {
	HasVoted bool       `bson:"hasVoted" json:"hasVoted"`
	Intent   VoteIntent `bson:"intent" json:"intent"`
	Winner   string     `bson:"winner" json:"winner"`
}

type Player struct {
	WalletAddress    string            `bson:"walletAddress" json:"walletAddress"`
	Vote             Vote              `bson:"vote" json:"vote"`
	HasStaked        bool              `bson:"hasStaked" json:"hasStaked"`
	StakeStatus      InteractionStatus `bson:"stakeStatus" json:"stakeStatus"`
	HasWithdrawn     bool              `bson:"hasWithdrawn" json:"hasWithdrawn"`
	WithdrawalStatus InteractionStatus `bson:"withdrawalStatus" json:"withdrawalStatus"`
}

type Wager struct {
	TransactionHash string            `bson:"transactionHash" json:"transactionHash"`
	SmartContractId string            `bson:"smart_contract_id" json:"smartContractId"`
	Creator         string            `bson:"creator" json:"creator"`
	CreationStatus  InteractionStatus `bson:"creationStatus" json:"creationStatus"`
	Name            string            `bson:"name" json:"name"`
	Category        string            `bson:"category" json:"category"`
	Description     string            `bson:"description" json:"description"`
	Location        string            `bson:"location" json:"location"`
	Stake           int64             `bson:"stake" json:"stake"`
	Currency        string            `bson:"currency" json:"currency"`
	Chain           string            `bson:"chain" json:"chain"`
	Participants    []Player          `bson:"participants" json:"participants"`
	Status          WagerStatus       `bson:"status" json:"status"`
	Winner          string            `bson:"winner" json:"winner"`
}

type PersistedWager struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt string        `bson:"createdAt" json:"createdAt"`
	UpdatedAt string        `bson:"updatedAt" json:"updatedAt"`
	Wager     `bson:",inline" json:",inline"`
}
