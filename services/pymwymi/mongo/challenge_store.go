package mongo

import (
	"context"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ChallengeStore interface {
	// need some sort of mongo connection here
	getChallengesForUser(walletAddress string, status pymwymi.ChallengeStatus, pageOpts pymwymi.PageOpts)
}

type Storage struct {
	c *mongo.Collection
}

func (s *Storage) NewChallengeStore(client *mongo.Client, dbName string) *Storage {
	return &Storage{
		c: client.Database(dbName).Collection("challenges"),
	}
}

func (s *Storage) CreateChallenge(ctx context.Context, challenge *pymwymi.Challenge) error {
	_, err := s.c.InsertOne(ctx, challenge)
	return err
}

// replaces the whole challenge so be wary of this
func (s *Storage) UpdateChallenge(ctx context.Context, challenge *pymwymi.PersistedChallenge) error {
	filter := bson.D{}
	objectID, err := bson.ObjectIDFromHex(challenge.ID)
	if err != nil {
		return err
	}
	filter = append(filter, bson.E{Key: "_id", Value: objectID})
	_, err = s.c.UpdateOne(ctx, bson.M{"transactionHash": challenge.TransactionHash}, bson.M{"$set": challenge})
	return err
}

// you can submit an empty walletAddress but not an empty status
func (s *Storage) getChallengesForUser(
	ctx context.Context, walletAddress string,
	status pymwymi.ChallengeStatus,
	pageOpts pymwymi.PageOpts,
) (*[]pymwymi.PersistedChallenge, error) {
	result := []pymwymi.PersistedChallenge{}
	filter := bson.D{}
	filter = append(filter, bson.E{Key: "status", Value: status})
	if walletAddress != "" {
		filter = append(filter, bson.E{Key: "walletAddress", Value: walletAddress})
	}
	options := setPageOpts(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, &result)
	return &result, nil
}

// if you pass 0 for page and limit it will return all
func setPageOpts(opts *options.FindOptionsBuilder, p pymwymi.PageOpts) *options.FindOptionsBuilder {
	if p.Page != 0 && p.Limit != 0 {
		opts.SetSkip((p.Page - 1) * p.Limit)
	}
	if p.Limit != 0 {
		opts.SetLimit(p.Limit)
	}
	return opts
}
