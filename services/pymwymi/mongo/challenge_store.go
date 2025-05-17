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

func (s *Storage) getChallengesForUser(
	ctx context.Context, walletAddress string,
	status pymwymi.ChallengeStatus,
	pageOpts pymwymi.PageOpts,
) (*[]pymwymi.Challenge, error) {
	// return challenges matching wallet Address and status
	result := []pymwymi.Challenge{}
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
