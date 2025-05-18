package challenge

import (
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
)

type ChallengeService interface {
	GetChallengesForUser(address string, status pymwymi.ChallengeStatus)
	CreateChallenge(challenge pymwymi.Challenge) error
	UpdateChallenge(challenge pymwymi.Challenge) error
	DeleteChallenge(id string) error
}

type Service struct {
	// need db connect here
	storage mongo.ChallengeStore
}

const PAGE_LIMIT = 50

func (s *Service) getChallengesForUser(walletAddress string, status pymwymi.ChallengeStatus, page int) {
	// do i need to paginate??? i think so
	// i should use middleware for this
	// and then access page via ctx?
	// challenges, err := s.storage.GetChallengesForUser(a)
}

func (s *Service) createChallenge(challenge pymwymi.Challenge) error {
	return s.storage.CreateChallenge(challenge)
}
