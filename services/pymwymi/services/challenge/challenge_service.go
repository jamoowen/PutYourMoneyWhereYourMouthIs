package challenge

import (
	"context"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
)

type Service struct {
	storage mongo.ChallengeStore
}

func NewChallengeService(storage mongo.ChallengeStore) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetChallengesForUser(ctx context.Context, status pymwymi.ChallengeStatus) ([]pymwymi.Challenge, error) {
	pageOpts := pymwymi.GetPageOptsFromCtx(ctx)
	walletAddress := pymwymi.GetUserFromCtx(ctx).WalletAddress
	persistedChallenges, err := s.storage.GetChallengesByStatus(ctx, walletAddress, status, pageOpts)
	if err != nil {
		return nil, err
	}
	// adapt
	challenges := make([]pymwymi.Challenge, len(persistedChallenges))
	for i, persistedChallenge := range persistedChallenges {
		challenges[i] = persistedChallenge.Challenge
	}
	return challenges, nil
}

func (s *Service) CreateChallenge(ctx context.Context, challenge pymwymi.NewChallengeDto) (pymwymi.Challenge, error) {
	// so we are passed all the stuff but we need to validate first
}

// func (s *Service) createChallenge(user) error {
// 	return s.storage.CreateChallenge(challenge)
// }
