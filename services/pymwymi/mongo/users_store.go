package mongo

import (
	"context"
	"fmt"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsersStorage struct {
	c *mongo.Collection
}

func NewUsersStore(client *mongo.Client, dbName string) *UsersStorage {
	return &UsersStorage{
		c: client.Database(dbName).Collection("users"),
	}
}

func (s *UsersStorage) CreateUser(ctx context.Context, user pymwymi.User) *pymwymi.Error {
	_, err := s.c.InsertOne(ctx, user)
	if err != nil {
		return pymwymi.Errorf(pymwymi.ErrInternal, "failed to insert new user: %s", err)
	}
	return nil
}

func (s *UsersStorage) GetUsersByWalletAddress(ctx context.Context, wallets []string) ([]pymwymi.PersistedUser, *pymwymi.Error) {
	var results []pymwymi.PersistedUser
	filter := bson.D{
		{Key: "walletAddress", Value: bson.D{
			{Key: "$in", Value: wallets},
		}},
	}
	cursor, err := s.c.Find(ctx, filter)
	if err != nil {
		return results, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get users (%v): %s", wallets, err.Error())
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, pymwymi.Errorf(pymwymi.ErrInternal, "failed to decode users (%v): %s", wallets, err.Error())
	}
	return results, nil
}

func (s *UsersStorage) GetUser(ctx context.Context, walletAddress string) (pymwymi.PersistedUser, *pymwymi.Error) {
	var result pymwymi.PersistedUser
	filter := bson.D{
		{Key: "walletAddress", Value: walletAddress},
	}
	err := s.c.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, pymwymi.Errorf(pymwymi.ErrUserNotFound, "user not found")
	}
	if err != nil {
		return result, pymwymi.Errorf(pymwymi.ErrUserNotFound, "failed to get user (%v): %s", walletAddress, err)
	}
	return result, nil
}

// this supports lots of fields but for now we just should be passing name as this is the only thing that makes sense
func (s *UsersStorage) UpdateUser(ctx context.Context, id string, fieldsToSet []pymwymi.FieldToSet) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{}
	for _, field := range fieldsToSet {
		update = append(update, bson.E{Key: field.Field, Value: field.Value})
	}
	_, err := s.c.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return fmt.Errorf("failed to update user (%v): %w", id, err)
	}
	return nil
}
