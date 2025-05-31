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

// once all players have staked, it changes to pending state
// if all the staked players have voted cancel, we should cancel
// this is a race condition. we should handle votes in a queue...
func (s *Service) SubmitVote(ctx context.Context, challengeID string) *pymwymi.Error {
	challenge, err := s.challengeStorage.GetChallengeByID(ctx, challengeID)
	if err != nil {
		return pymwymi.Errorf(err.Code, "could not submit cancel vote: %v", err)
	}
	if challenge.Status != pymwymi.StateCreated && challenge.Status != pymwymi.StatePending {
		return pymwymi.Errorf(pymwymi.ErrBadInput, "you can only vote on a challenge that is created or pending")
	}
	return nil
}

// we just need to check that all members who have staked have voted
// if votes are unanimous we can cancel the challenge
func (s *Service) SubmitCancelVote(ctx context.Context, challenge *pymwymi.PersistedChallenge) *pymwymi.Error {
	user := pymwymi.GetUserFromCtx(ctx)
	unanimousVotes := true
	for _, participant := range challenge.Participants {
		if !participant.HasStaked {
			continue
		}
		if participant.WalletAddress == user.WalletAddress {
			participant.Vote.Intention = pymwymi.VoteCancel
			participant.Vote.Winner = ""
		} else if participant.Vote.Intention == pymwymi.VoteWinner {
			unanimousVotes = false
		}
	}
	if unanimousVotes == true {
		challenge.Status = pymwymi.StateCancelled
	}
	// @TODO we should now update the mongo record but also make sure to call the smart contract
	// would be good to have something like kafka here since we absolutely need to call the sc
	// perhaps we need another state to indicate that its in the process of being cancelled
	return nil
}

// func (s *Service) createChallenge(user) error {
// 	return s.storage.CreateChallenge(challenge)
// }
