package mongo

import (
	"context"
	"fmt"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type WagerStorage struct {
	c *mongo.Collection
}

func NewWagerStore(client *mongo.Client, dbName string) *WagerStorage {
	return &WagerStorage{
		c: client.Database(dbName).Collection("wagers"),
	}
}

func (s *WagerStorage) CreateWager(ctx context.Context, wager *pymwymi.Wager) *pymwymi.Error {
	_, err := s.c.InsertOne(ctx, wager)
	if err != nil {
		return pymwymi.Errorf(pymwymi.ErrInternal, "failed to create wager: %s", err.Error())
	}
	return nil
}

// this has race conditions and is important to only allow a single vote to make it through
func (s *WagerStorage) UpdateWagerWithVote(ctx context.Context, id string, wager pymwymi.Wager) *pymwymi.Error {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return pymwymi.Errorf(pymwymi.ErrBadInput, "invalid wager id (%v): %v", id, err)
	}
	validStatuses := []pymwymi.WagerStatus{pymwymi.StateCreated, pymwymi.StatePending}
	filter := bson.D{
		{Key: "_id", Value: objectId},
		{Key: "status", Value: bson.D{{Key: "$in", Value: validStatuses}}},
	}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: wager.Status},
		bson.E{Key: "participants", Value: wager.Participants},
	}}}
	result, err := s.c.UpdateOne(ctx, filter, update)
	if err != nil {
		return pymwymi.Errorf(pymwymi.ErrInternal, "failed to update wager (%v): %v", id, err)
	}
	// at this point we have already verified the wager exists
	// if there are no wagers found - its the status part of the filter that isnt matching
	// ie its not in pending or created state and therefore voting is over
	if result.MatchedCount == 0 {
		return pymwymi.Errorf(pymwymi.ErrVotingFinished, "voting in this wager no longer possible")
	}
	return nil
}

func (s *WagerStorage) UpdateWager(ctx context.Context, id string, fieldsToSet []pymwymi.FieldToSet) error {
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
		return fmt.Errorf("failed to update wager (%v): %w", id, err)
	}
	return nil
}

// we need to handle not found and a genuine db error differently
func (s *WagerStorage) GetWagerByID(ctx context.Context, id string) (pymwymi.PersistedWager, *pymwymi.Error) {
	var wager pymwymi.PersistedWager
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return wager, pymwymi.Errorf(pymwymi.ErrBadInput, "invalid wager id (%v): %v", id, err)
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	err = s.c.FindOne(ctx, filter).Decode(&wager)
	if err == mongo.ErrNoDocuments {
		return wager, pymwymi.Errorf(pymwymi.ErrWagerNotFound, "wager (%v) not found", id)
	}
	if err != nil {
		return wager, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get wager (%v): %v", id, err)
	}
	return wager, nil
}

// you can submit an empty walletAddress but not an empty status
func (s *WagerStorage) GetWagersByStatus(
	ctx context.Context,
	walletAddress string,
	status pymwymi.WagerStatus,
	pageOpts pymwymi.PageOpts,
) ([]pymwymi.PersistedWager, error) {
	result := []pymwymi.PersistedWager{}
	filter := bson.D{bson.E{Key: "status", Value: status}}
	if walletAddress != "" {
		filter = append(filter, bson.E{Key: "walletAddress", Value: walletAddress})
	}
	options := setPageOptions(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return result, fmt.Errorf("failed to get wagers: %w", err)
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
