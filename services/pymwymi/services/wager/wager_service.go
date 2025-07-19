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

func (s *Service) GetWagerForParticipant(ctx context.Context, id string, walletAddress string) (*pymwymi.PersistedWager, *pymwymi.Error) {
	persistedWager, err := s.wagerStorage.GetWagerByID(ctx, id)
	if err != nil {
		return nil, pymwymi.Errorf(err.Code, "failed to get wager by id: %s", err.Message)
	}
	isParticipant := false
	for _, participant := range persistedWager.Participants {
		if participant.WalletAddress == walletAddress {
			isParticipant = true
			break
		}
	}
	if isParticipant == false {
		return nil, pymwymi.Errorf(pymwymi.ErrNotParticipant, "not a participant in this wager")
	}
	return persistedWager, nil
}

func (s *Service) GetWagersForUser(
	ctx context.Context,
	creator string,
	pageOpts *pymwymi.PageOpts,
	status pymwymi.WagerStatus,
	walletAddress string,
	winner string,
) ([]pymwymi.PersistedWager, *pymwymi.Error) {
	var wagers []pymwymi.PersistedWager
	var err *pymwymi.Error

	switch status {
	case pymwymi.StateCreated:
		wagers, err = s.wagerStorage.GetCreatedWagers(ctx, walletAddress, creator == "true", pageOpts)
	case pymwymi.StatePending:
		wagers, err = s.wagerStorage.GetPendingWagers(ctx, walletAddress, pageOpts)
	case pymwymi.StateCompleted:
		wagers, err = s.wagerStorage.GetClaimableWagers(ctx, walletAddress, pageOpts)
	case pymwymi.StateClaimed:
		wagers, err = s.wagerStorage.GetPastWagers(ctx, walletAddress, pageOpts)
	}

	if err != nil {
		return nil, pymwymi.Errorf(err.Code, "failed to get wagers by status: %s", err.Message)
	}
	return wagers, nil
}

// @DEV I thought we should store the pymwymi users on the record but just walletAddress is probably fine
func (s *Service) CreateWager(ctx context.Context,
	category string,
	currency string,
	description string,
	location string,
	name string,
	participantsWalletAddresses []string,
	stake int64,
	transactionHash string,
) (pymwymi.Wager, *pymwymi.Error) {
	creator := pymwymi.GetUserFromCtx(ctx).WalletAddress

	participants := make([]pymwymi.Player, len(participantsWalletAddresses)+1)
	for i, address := range participantsWalletAddresses {
		participants[i] = pymwymi.Player{
			WalletAddress: address,
		}
	}
	participants[len(participants)-1] = pymwymi.Player{WalletAddress: creator}

	status := pymwymi.StateCreated
	creationStatus := pymwymi.InteractionStatePending

	wager := pymwymi.Wager{
		TransactionHash: transactionHash,
		Creator:         creator,
		CreationStatus:  creationStatus,
		Name:            name,
		Category:        category,
		Description:     description,
		Location:        location,
		Stake:           stake,
		Currency:        currency,
		Participants:    participants,
		Status:          status,
	}

	err := s.wagerStorage.CreateWager(ctx, wager)
	if err != nil {
		return pymwymi.Wager{}, pymwymi.Errorf(err.Code, "failed to create wager: %s", err.Message)
	}
	return wager, nil
}

// we just need to check that all members who have staked have voted
// if nobody accepts then the creator should be allowed to cancel
// if votes are unanimous we can cancel the wager
// we are handling race condtions in the mongo operation
func (s *Service) SubmitVote(ctx context.Context, user pymwymi.User, wager *pymwymi.PersistedWager, vote pymwymi.Vote) *pymwymi.Error {
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
