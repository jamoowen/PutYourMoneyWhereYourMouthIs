package mongo

import "github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/challengeapi"

type ChallengeStore interface {
	GetUnacceptedChallenges(wallets []string) ([]challengeapi.Challenge, error)
	GetCreatedChallenges(wallets []string) ([]challengeapi.Challenge, error)
	GetDisputedChallenges(wallets []string) ([]challengeapi.Challenge, error)
	GetClaimableChallenges(wallets []string) ([]challengeapi.Challenge, error)
	GetChallengeById(id string) (challengeapi.Challenge, error)
	CreateChallenge(challenge challengeapi.Challenge) error
	UpdateChallenge(challenge challengeapi.Challenge) error
	DeleteChallenge(id string) error
}
