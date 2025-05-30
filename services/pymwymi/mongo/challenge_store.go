package mongo

import (
	"context"
	"fmt"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ChallengeStorage struct {
	c *mongo.Collection
}

func NewChallengeStore(client *mongo.Client, dbName string) *ChallengeStorage {
	return &ChallengeStorage{
		c: client.Database(dbName).Collection("challenges"),
	}
}

func (s *ChallengeStorage) CreateChallenge(ctx context.Context, challenge *pymwymi.Challenge) error {
	_, err := s.c.InsertOne(ctx, challenge)
	if err != nil {
		return fmt.Errorf("failed to create challenge: %w", err)
	}
	return nil
}

func (s *ChallengeStorage) UpdateChallenge(ctx context.Context, id string, fieldsToSet []pymwymi.FieldToSet) error {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	setDoc := bson.D{}
	for _, field := range fieldsToSet {
		setDoc = append(setDoc, bson.E{Key: field.Field, Value: field.Value})
	}
	_, err = s.c.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: setDoc}})
	if err != nil {
		return fmt.Errorf("failed to update challenge (%v): %w", id, err)
	}
	return nil
}

// we need to handle not found and a genuine db error differently
func (s *ChallengeStorage) GetChallengeByID(ctx context.Context, id string) (*pymwymi.Challenge, error) {
	var challenge pymwymi.Challenge
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge (%v): %w", id, err)
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	err = s.c.FindOne(ctx, filter).Decode(&challenge)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge (%v): %w", id, err)
	}
	return &challenge, nil
}

// you can submit an empty walletAddress but not an empty status
func (s *ChallengeStorage) GetChallengesByStatus(
	ctx context.Context,
	walletAddress string,
	status pymwymi.ChallengeStatus,
	pageOpts pymwymi.PageOpts,
) ([]pymwymi.PersistedChallenge, error) {
	result := []pymwymi.PersistedChallenge{}
	filter := bson.D{bson.E{Key: "status", Value: status}}
	if walletAddress != "" {
		filter = append(filter, bson.E{Key: "walletAddress", Value: walletAddress})
	}
	options := setPageOptions(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return result, fmt.Errorf("failed to get challenges: %w", err)
	}
	cursor.All(ctx, &result)
	return result, nil
}

// if you pass 0 for page and limit it will return all
func setPageOptions(opts *options.FindOptionsBuilder, p pymwymi.PageOpts) *options.FindOptionsBuilder {
	if p.Page != 0 && p.Limit != 0 {
		opts.SetSkip((p.Page - 1) * p.Limit)
	}
	if p.Limit != 0 {
		opts.SetLimit(p.Limit)
	}
	return opts
}
