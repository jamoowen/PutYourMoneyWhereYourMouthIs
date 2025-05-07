package mongo

import "github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"

type ChallengeStore interface {
	GetUnacceptedChallenges(wallets []string) ([]pymwymi.Challenge, error)
	GetCreatedChallenges(wallets []string) ([]pymwymi.Challenge, error)
	GetDisputedChallenges(wallets []string) ([]pymwymi.Challenge, error)
	GetClaimableChallenges(wallets []string) ([]pymwymi.Challenge, error)
	GetChallengeById(id string) (pymwymi.Challenge, error)
	CreateChallenge(challenge pymwymi.Challenge) error
	UpdateChallenge(challenge pymwymi.Challenge) error
	DeleteChallenge(id string) error
}
