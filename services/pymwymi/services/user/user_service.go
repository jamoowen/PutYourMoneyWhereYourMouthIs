package user

import (
	"context"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
)

type Service struct {
	storage *mongo.UsersStorage
}

func NewUserService(storage *mongo.UsersStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateUser(ctx context.Context, walletAddress string) (pymwymi.User, *pymwymi.Error) {
	newUser := pymwymi.User{
		WalletAddress: walletAddress,
		Name:          pymwymi.DEFAULT_USER_NAME,
	}
	err := s.storage.CreateUser(ctx, newUser)
	if err != nil {
		return pymwymi.User{}, pymwymi.Errorf(pymwymi.ErrInternal, "failed to create user: %s", err.Error())
	}
	return newUser, nil
}

func (s *Service) GetUser(ctx context.Context, walletAddress string) (pymwymi.User, *pymwymi.Error) {
	persistedUser, err := s.storage.GetUser(ctx, walletAddress)
	if err != nil {
		return pymwymi.User{}, pymwymi.Errorf(err.Code, "failed to get user: %s", err.Error())
	}

	user := pymwymi.User{
		WalletAddress: persistedUser.WalletAddress,
		Name:          persistedUser.Name,
	}
	return user, nil
}

func (s *Service) GetUsers(ctx context.Context, wallets []string) ([]pymwymi.User, error) {
	var users []pymwymi.User
	persistedUsers, err := s.storage.GetUsersByWalletAddress(ctx, wallets)
	if err != nil {
		return users, err
	}
	for _, persistedUser := range persistedUsers {
		user := pymwymi.User{
			WalletAddress: persistedUser.WalletAddress,
			Name:          persistedUser.Name,
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Service) UpdateName(ctx context.Context, newName string, walletAddress string) (pymwymi.User, *pymwymi.Error) {
	updatedUser, err := s.storage.UpdateName(ctx, newName, walletAddress)
	if err != nil {
		return pymwymi.User{}, pymwymi.Errorf(err.Code, "failed to update user name: %s", err.Error())
	}
	user := pymwymi.User{
		WalletAddress: updatedUser.WalletAddress,
		Name:          updatedUser.Name,
	}

	return user, nil
}
