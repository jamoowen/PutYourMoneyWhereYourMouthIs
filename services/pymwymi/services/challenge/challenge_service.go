package challenge

import (
	"context"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
)

type ChallengeService interface {
	GetChallengesForUser(ctx context.Context, status pymwymi.ChallengeStatus, page int64) ([]pymwymi.Challenge, error)
}

type Service struct {
	storage mongo.ChallengeStore
}

const PAGE_LIMIT = 50

func (s *Service) getChallengesForUser(ctx context.Context, status pymwymi.ChallengeStatus, page int64) ([]pymwymi.Challenge, error) {
	pageOpts := pymwymi.PageOpts{
		Page:  page,
		Limit: PAGE_LIMIT,
	}
	persistedChallenges, err := s.storage.GetChallengesByStatus(pymwymi.GetUserFromCtx(ctx).WalletAddress, status, pageOpts)
	if err != nil {
		return nil, err
	}
	// convert (remove _id and createdAt)
	challenges := make([]pymwymi.Challenge, len(persistedChallenges))
	for i, persistedChallenge := range persistedChallenges {
		challenges[i] = persistedChallenge.Challenge
	}
	return challenges, nil
}

// func (s *Service) createChallenge(user) error {
// 	return s.storage.CreateChallenge(challenge)
// }
