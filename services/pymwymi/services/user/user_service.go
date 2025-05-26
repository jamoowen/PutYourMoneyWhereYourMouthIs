package user

import (
	"context"
	"fmt"

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

func (s *Service) CreateUser(ctx context.Context, walletAddress string) error {
	persistedUsers, err := s.storage.GetUsers(ctx, []string{walletAddress})
	if err != nil {
		return fmt.Errorf("failed to get create user: %w", err)
	}
	if len(persistedUsers) > 0 {
		return nil
	}
	return s.storage.CreateUser(ctx, &pymwymi.User{
		WalletAddress: walletAddress,
		Name:          pymwymi.DEFAULT_USER_NAME,
	})
}

func (s *Service) GetUsers(ctx context.Context, wallets []string) ([]pymwymi.User, error) {
	var users []pymwymi.User
	persistedUsers, err := s.storage.GetUsers(ctx, wallets)
	if err != nil {
		return users, err
	}
	for _, persistedUser := range persistedUsers {
		users = append(users, persistedUser.User)
	}
	return users, nil
}

func (s *Service) UpdateUserName(ctx context.Context, id string, name string) error {
	fieldsToSet := []pymwymi.FieldToSet{
		{
			Field: "name",
			Value: name,
		},
	}
	return s.storage.UpdateUser(ctx, id, fieldsToSet)
}
