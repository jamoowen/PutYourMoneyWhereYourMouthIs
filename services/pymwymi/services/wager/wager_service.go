package wager

import (
	"context"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
)

type Service struct {
	wagerStorage *mongo.WagerStorage
	userStorage  *mongo.UsersStorage
}

func NewWagerService(wagerStorage *mongo.WagerStorage, userStorage *mongo.UsersStorage) *Service {
	return &Service{
		wagerStorage: wagerStorage,
		userStorage:  userStorage,
	}
}

func (s *Service) GetWager(ctx context.Context, id string) (*pymwymi.Wager, error) {
	persistedWager, err := s.wagerStorage.GetWagerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &persistedWager.Wager, nil
}

func (s *Service) GetWagerForParticipant(ctx context.Context, id string, walletAddress string) (pymwymi.Wager, *pymwymi.Error) {
	persistedWager, err := s.wagerStorage.GetWagerByID(ctx, id)
	if err != nil {
		return pymwymi.Wager{}, err
	}
	isParticipant := false
	for _, participant := range persistedWager.Participants {
		if participant.WalletAddress == walletAddress {
			isParticipant = true
			break
		}
	}
	if isParticipant == false {
		return pymwymi.Wager{}, pymwymi.Errorf(pymwymi.ErrNotParticipant, "not a participant in this wager")
	}
	return persistedWager.Wager, nil
}

func (s *Service) GetWagersForUser(ctx context.Context, status pymwymi.WagerStatus) ([]pymwymi.Wager, error) {
	pageOpts := pymwymi.GetPageOptsFromCtx(ctx)
	walletAddress := pymwymi.GetUserFromCtx(ctx).WalletAddress
	persistedWagers, err := s.wagerStorage.GetWagersByStatus(ctx, walletAddress, status, pageOpts)
	if err != nil {
		return nil, err
	}
	// adapt
	wagers := make([]pymwymi.Wager, len(persistedWagers))
	for i, persistedWager := range persistedWagers {
		wagers[i] = persistedWager.Wager
	}
	return wagers, nil
}

// @DEV I thought we should store the pymwymi users on the record but just walletAddress is probably fine
func (s *Service) CreateWager(ctx context.Context,
	transactionHash,
	name,
	category,
	description,
	location,
	stake,
	currency string,
	participantsWalletAddresses []string,
) (pymwymi.Wager, *pymwymi.Error) {
	creator := pymwymi.GetUserFromCtx(ctx).WalletAddress

	participants := make([]pymwymi.Player, len(participantsWalletAddresses)+1)
	for i, address := range participantsWalletAddresses {
		participants[i] = pymwymi.Player{
			WalletAddress: address,
		}
	}
	participants[len(participants)] = pymwymi.Player{
		WalletAddress: creator,
	}

	status := pymwymi.StateCreated
	newWager := pymwymi.Wager{
		TransactionHash: transactionHash,
		Creator:         creator,
		Name:            name,
		Category:        category,
		Description:     description,
		Location:        location,
		Stake:           stake,
		Currency:        currency,
		Participants:    participants,
		Status:          status,
	}
	err := s.wagerStorage.CreateWager(ctx, &newWager)
	if err != nil {
		return pymwymi.Wager{}, pymwymi.Errorf(err.Code, "failed to create wager: %s", err.Message)
	}
	return newWager, nil
}

// // once all players have staked, it changes to pending state
// // if all the staked players have voted cancel, we should cancel
// // this is a race condition. we should handle votes in a queue...
// func (s *Service) SubmitVote(ctx context.Context, wagerID string) *pymwymi.Error {
// 	wager, err := s.wagerStorage.GetWagerByID(ctx, wagerID)
// 	if err != nil {
// 		return pymwymi.Errorf(err.Code, "could not submit cancel vote: %v", err)
// 	}
// 	if wager.Status != pymwymi.StateCreated && wager.Status != pymwymi.StatePending {
// 		return pymwymi.Errorf(pymwymi.ErrBadInput, "you can only vote on a wager that is created or pending")
// 	}
// 	return nil
// }

// we just need to check that all members who have staked have voted
// if nobody accepts then the creator should be allowed to cancel
// if votes are unanimous we can cancel the wager
// we are handling race condtions in the mongo operation
func (s *Service) SubmitVote(ctx context.Context, user pymwymi.User, wager pymwymi.Wager, vote pymwymi.Vote) *pymwymi.Error {
	unanimousVotes := true
	for i, participant := range wager.Participants {
		if !participant.HasStaked {
			continue
		}
		if participant.WalletAddress == user.WalletAddress {
			wager.Participants[i].Vote = vote
			continue
		}
		if !participant.Vote.HasVoted || participant.Vote.Intent != vote.Intent {
			unanimousVotes = false
			continue
		}
	}
	if unanimousVotes == true {
		wager.Status = pymwymi.StateCancelled
	}
	err := s.wagerStorage.UpdateWagerWithVote(ctx, wager.ID, wager)
	if err != nil {
		return pymwymi.Errorf(err.Code, "failed to submit cancel vote: %v", err)
	}
	return nil
}

// func (s *Service) createWager(user) error {
// 	return s.storage.CreateWager(wager)
// }
