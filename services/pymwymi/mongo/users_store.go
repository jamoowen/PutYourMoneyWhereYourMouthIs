package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
	now := pymwymi.IsoNow()
	persistedUser := pymwymi.PersistedUser{
		CreatedAt: now,
		UpdatedAt: now,
		User:      user,
	}
	b, _ := bson.MarshalExtJSON(persistedUser, false, false)
	fmt.Println(string(b))
	_, err := s.c.InsertOne(ctx, persistedUser)
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
		return results, pymwymi.Errorf(pymwymi.ErrInternal, "%s", err.Error())
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, pymwymi.Errorf(pymwymi.ErrInternal, "failed to decode users %s", err.Error())
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
		return result, pymwymi.Errorf(pymwymi.ErrInternal, "%s", err.Error())
	}
	return result, nil
}

func (s *UsersStorage) UpdateName(ctx context.Context, name, walletAddress string) (pymwymi.PersistedUser, *pymwymi.Error) {
	now := time.Now().UTC().Format(time.RFC3339)
	filter := bson.D{
		{Key: "walletAddress", Value: walletAddress},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "updatedAt", Value: now},
		}},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedUser pymwymi.PersistedUser
	err := s.c.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedUser)
	if err == mongo.ErrNoDocuments {
		return pymwymi.PersistedUser{}, pymwymi.Errorf(pymwymi.ErrUserNotFound, "user not found")
	}
	if err != nil {
		return pymwymi.PersistedUser{}, pymwymi.Errorf(pymwymi.ErrInternal, "%s", err)
	}
	log.Printf("updated user: %v", updatedUser)
	return updatedUser, nil
}
