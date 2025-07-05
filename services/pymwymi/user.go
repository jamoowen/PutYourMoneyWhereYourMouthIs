package pymwymi

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

const DEFAULT_USER_NAME = "Unknown"

type User struct {
	Name          string `bson:"name" json:"name"`
	WalletAddress string `bson:"walletAddress" json:"walletAddress"`
}

type PersistedUser struct {
	ID        bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatedAt string        `json:"createdAt" bson:"createdAt"`
	UpdatedAt string        `json:"updatedAt" bson:"updatedAt"`
	User      `bson:",inline" json:",inline"`
}
