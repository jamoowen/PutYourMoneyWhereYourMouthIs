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
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt string        `bson:"createdAt" json:"createdAt"`
	UpdatedAt string        `bson:"updatedAt" json:"updatedAt"`
	User      `bson:",inline" json:",inline"`
}
