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

func (s *Service) GetChallenge(ctx context.Context, id string) (*pymwymi.Challenge, error) {
	persistedChallenge, err := s.challengeStorage.GetChallengeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &persistedChallenge.Challenge, nil
}

func (s *Service) GetChallengeForParticipant(ctx context.Context, id string, walletAddress string) (*pymwymi.Challenge, error) {
	persistedChallenge, err := s.challengeStorage.GetChallengeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	isParticipant := false
	for _, participant := range persistedChallenge.Participants {
		if participant.WalletAddress == walletAddress {
			isParticipant = true
			break
		}
	}
	if isParticipant == false {
		return nil, pymwymi.Errorf(pymwymi.ErrNotParticipant, "not a participant in this challenge")
	}
	return &persistedChallenge.Challenge, nil
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
			WalletAddress: participant.WalletAddress,
		})
	}
	newChallenge := pymwymi.Challenge{
		TransactionHash: challenge.TransactionHash,
		Creator:         challenge.Creator,
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

// // once all players have staked, it changes to pending state
// // if all the staked players have voted cancel, we should cancel
// // this is a race condition. we should handle votes in a queue...
// func (s *Service) SubmitVote(ctx context.Context, challengeID string) *pymwymi.Error {
// 	challenge, err := s.challengeStorage.GetChallengeByID(ctx, challengeID)
// 	if err != nil {
// 		return pymwymi.Errorf(err.Code, "could not submit cancel vote: %v", err)
// 	}
// 	if challenge.Status != pymwymi.StateCreated && challenge.Status != pymwymi.StatePending {
// 		return pymwymi.Errorf(pymwymi.ErrBadInput, "you can only vote on a challenge that is created or pending")
// 	}
// 	return nil
// }

// we just need to check that all members who have staked have voted
// if nobody accepts then the creator should be allowed to cancel
// if votes are unanimous we can cancel the challenge
// we are handling race condtions in the mongo operation
func (s *Service) SubmitVote(ctx context.Context, user pymwymi.User, challenge *pymwymi.Challenge, vote pymwymi.Vote) error {
	unanimousVotes := true
	for i, participant := range challenge.Participants {
		if !participant.HasStaked {
			continue
		}
		if participant.WalletAddress == user.WalletAddress {
			challenge.Participants[i].Vote = vote
			continue
		}
		if !participant.Vote.HasVoted || participant.Vote.Intent != vote.Intent {
			unanimousVotes = false
			continue
		}
	}
	if unanimousVotes == true {
		challenge.Status = pymwymi.StateCancelled
	}
	err := s.challengeStorage.UpdateChallengeWithVote(ctx, challenge.ID, challenge)
	if err != nil {
		return pymwymi.Errorf(err.Code, "failed to submit cancel vote: %v", err)
	}
	return nil
}

// func (s *Service) createChallenge(user) error {
// 	return s.storage.CreateChallenge(challenge)
// }
