package challenge

import (
	"context"
	"fmt"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
)

type Service struct {
	challengeStorage *mongo.ChallengeStorage
	userStorage      *mongo.UsersStorage
}

func NewChallengeService(challengeStorage *mongo.ChallengeStorage, userStorage *mongo.UsersStorage) *Service {
	return &Service{
		challengeStorage: challengeStorage,
		userStorage:      userStorage,
	}
}

func (s *Service) GetChallengesForUser(ctx context.Context, status pymwymi.ChallengeStatus) ([]pymwymi.Challenge, error) {
	pageOpts := pymwymi.GetPageOptsFromCtx(ctx)
	walletAddress := pymwymi.GetUserFromCtx(ctx).WalletAddress
	persistedChallenges, err := s.challengeStorage.GetChallengesByStatus(ctx, walletAddress, status, pageOpts)
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

func (s *Service) CreateChallenge(ctx context.Context, challenge pymwymi.NewChallengeDto) (*pymwymi.Challenge, error) {
	// so we are passed all the stuff but we need to validate first
	users, err := s.userStorage.GetUsersByWalletAddress(ctx, challenge.ParticipantsAddresses)
	if err != nil {
		return nil, err
	}
	if len(users) != len(challenge.ParticipantsAddresses) {
		return nil, fmt.Errorf("all participants must have signed into pymwymi before creating a challenge")
	}
	var participants []pymwymi.Player
	for _, participant := range users {
		participants = append(participants, pymwymi.Player{
			WalletAddress: pymwymi.WalletAddress(participant.WalletAddress),
		})
	}
	newChallenge := pymwymi.Challenge{
		TransactionHash: challenge.TransactionHash,
		Creator:         pymwymi.WalletAddress(challenge.Creator),
		Name:            challenge.Name,
		Category:        challenge.Category,
		Description:     challenge.Description,
		Location:        challenge.Location,
		Stake:           challenge.Stake,
		Currency:        challenge.Currency,
		Participants:    participants,
		Status:          pymwymi.StatePending,
	}
	err = s.challengeStorage.CreateChallenge(ctx, &newChallenge)
	if err != nil {
		return nil, fmt.Errorf("could not create challenge: %w", err)
	}
	return &newChallenge, nil
}

func (s *Service) SubmitUserVote(ctx context.Context, vote pymwymi.VoteDTO) error {
	challenge, err := s.challengeStorage.GetChallengeByID(ctx, vote.ChallengeId)
	if err != nil {
		return
	}
	if challenge == nil {
		return fmt.Errorf("challenge not found")
	}
	if challenge.Status != pymwymi.StatePending {
		return fmt.Errorf("challenge is not in pending state")
	}
	return nil
}

// func (s *Service) createChallenge(user) error {
// 	return s.storage.CreateChallenge(challenge)
// }
