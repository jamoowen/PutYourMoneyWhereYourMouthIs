package services

import (
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/challengeapi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/challengeapi/mongo"
)

type ChallengeService struct {
	store mongo.ChallengeStore
}

func CreateChallenge() error {
}
