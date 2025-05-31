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

type VoteIntention string

const (
	VotePending VoteIntention = "pending"
	VoteCancel  VoteIntention = "cancel"
	VoteWinner  VoteIntention = "winner"
)

type Vote struct {
	Intention VoteIntention `bson:"intention" json:"intention"`
	Winner    string        `bson:"winner" json:"winner"`
}

type VoteDTO struct {
	ChallengeId string `json:"challengeId"`
	Vote        Vote   `json:"vote"`
}

type Player struct {
	WalletAddress string `bson:"walletAddress" json:"walletAddress"`
	Vote          Vote   `bson:"vote" json:"vote"`
	HasStaked     bool   `bson:"hasStaked" json:"hasStaked"`
	HasWithdrawn  bool   `bson:"hasWithdrawn" json:"hasWithdrawn"`
}

type NewChallengeDto struct {
	ID                    string   `json:"id"`
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
	ID              string          `bson:"_id" json:"_id"`
	TransactionHash string          `bson:"transactionHash" json:"transactionHash"`
	Creator         string          `bson:"creator" json:"creator"`
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
	CreatedAt int64 `json:"createdAt" bson:"createdAt"`
	Challenge
}
