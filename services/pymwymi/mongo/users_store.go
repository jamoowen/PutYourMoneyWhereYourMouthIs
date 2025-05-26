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

func (s *UsersStorage) CreateUser(ctx context.Context, user *pymwymi.User) error {
	_, err := s.c.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *UsersStorage) GetUsers(ctx context.Context, wallets []string) ([]pymwymi.PersistedUser, error) {
	var results []pymwymi.PersistedUser
	filter := bson.D{
		{Key: "walletAddress", Value: bson.D{
			{Key: "$in", Value: wallets},
		}},
	}
	cursor, err := s.c.Find(ctx, filter)
	if err != nil {
		return results, fmt.Errorf("failed to get users (%v): %w", wallets, err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, fmt.Errorf("failed to parse fetched users (%v): %w", wallets, err)
	}
	return results, nil
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
