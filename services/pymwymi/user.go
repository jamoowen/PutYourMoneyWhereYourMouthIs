package pymwymi

const DEFAULT_USER_NAME = "Unknown"

type User struct {
	Name          string `bson:"name" json:"name"`
	WalletAddress string `bson:"walletAddress" json:"walletAddress"`
}

type PersistedUser struct {
	ID        string `json:"_id" bson:"_id"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	User
}
