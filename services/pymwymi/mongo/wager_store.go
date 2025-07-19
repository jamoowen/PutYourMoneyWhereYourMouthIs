package mongo

import (
	"context"
	"log"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type WagerStorage struct {
	c *mongo.Collection
}

const CollectionName = "wagers"

func NewWagerStore(client *mongo.Client, dbName string) *WagerStorage {
	return &WagerStorage{
		c: client.Database(dbName).Collection(CollectionName),
	}
}

func (s *WagerStorage) CreateWager(ctx context.Context, wager pymwymi.Wager) *pymwymi.Error {
	now := pymwymi.IsoNow()
	persistedWager := pymwymi.PersistedWager{
		CreatedAt: now,
		UpdatedAt: now,
		Wager:     wager,
	}
	_, err := s.c.InsertOne(ctx, persistedWager)
	if err != nil {
		return pymwymi.Errorf(pymwymi.ErrInternal, "failed to create wager: %s", err.Error())
	}
	return nil
}

// this has race conditions and is important to only allow a single vote to make it through
// we are handling race conditions with channels and a queue
func (s *WagerStorage) UpdateWagerWithVote(ctx context.Context, id bson.ObjectID, wager *pymwymi.PersistedWager) *pymwymi.Error {
	validStatuses := []pymwymi.WagerStatus{pymwymi.StateCreated, pymwymi.StatePending}
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "status", Value: bson.D{{Key: "$in", Value: validStatuses}}},
	}
	update := bson.M{
		"$set": bson.M{
			"status":       wager.Status,
			"participants": wager.Participants,
		},
	}
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

// we need to handle not found and a genuine db error differently
func (s *WagerStorage) GetWagerByID(ctx context.Context, id string) (*pymwymi.PersistedWager, *pymwymi.Error) {
	var wager pymwymi.PersistedWager
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, pymwymi.Errorf(pymwymi.ErrBadInput, "invalid wager id (%v): %v", id, err)
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	err = s.c.FindOne(ctx, filter).Decode(&wager)
	if err == mongo.ErrNoDocuments {
		return nil, pymwymi.Errorf(pymwymi.ErrWagerNotFound, "wager (%v) not found", id)
	}
	if err != nil {
		return nil, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get wager (%v): %v", id, err)
	}
	return &wager, nil
}

func (s *WagerStorage) GetPastWagers(
	ctx context.Context,
	walletAddress string,
	pageOpts *pymwymi.PageOpts,
) ([]pymwymi.PersistedWager, *pymwymi.Error) {
	result := []pymwymi.PersistedWager{}
	filter := bson.D{
		{Key: "participants.walletAddress", Value: walletAddress},
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "status", Value: pymwymi.StateClaimed}},
			bson.D{{Key: "status", Value: pymwymi.StateCancelled}},
			bson.D{
				{Key: "status", Value: pymwymi.StateCompleted},
				{Key: "winner", Value: bson.D{{Key: "$ne", Value: walletAddress}}},
			},
		}},
	}
	options := setPageOptions(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return result, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get past wagers: %s", err.Error())
	}
	for cursor.Next(ctx) {
		var w pymwymi.PersistedWager
		if err := cursor.Decode(&w); err != nil {
			log.Printf("Failed to decode: %v\n", err)
			continue
		}
		result = append(result, w)
	}
	return result, nil
}

func (s *WagerStorage) GetClaimableWagers(
	ctx context.Context,
	walletAddress string,
	pageOpts *pymwymi.PageOpts,
) ([]pymwymi.PersistedWager, *pymwymi.Error) {
	result := []pymwymi.PersistedWager{}
	filter := bson.D{
		bson.E{Key: "status", Value: pymwymi.StateCreated},
		bson.E{Key: "participants.walletAddress", Value: walletAddress},
		bson.E{Key: "winner", Value: walletAddress},
	}
	options := setPageOptions(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return result, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get claimable wagers: %s", err.Error())
	}
	for cursor.Next(ctx) {
		var w pymwymi.PersistedWager
		if err := cursor.Decode(&w); err != nil {
			log.Printf("Failed to decode: %v\n", err)
			continue
		}
		result = append(result, w)
	}
	return result, nil
}

// you can submit an empty walletAddress but not an empty status
func (s *WagerStorage) GetCreatedWagers(
	ctx context.Context,
	walletAddress string,
	creator bool,
	pageOpts *pymwymi.PageOpts,
) ([]pymwymi.PersistedWager, *pymwymi.Error) {
	result := []pymwymi.PersistedWager{}
	filter := bson.D{
		bson.E{Key: "status", Value: pymwymi.StateCreated},
		bson.E{Key: "participants.walletAddress", Value: walletAddress},
	}
	if creator {
		filter = append(filter, bson.E{Key: "creator", Value: walletAddress})
	}
	options := setPageOptions(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return result, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get created wagers: %s", err.Error())
	}
	for cursor.Next(ctx) {
		var w pymwymi.PersistedWager
		if err := cursor.Decode(&w); err != nil {
			log.Printf("Failed to decode: %v\n", err)
			continue
		}
		result = append(result, w)
	}
	return result, nil
}

// you can submit an empty walletAddress but not an empty status
func (s *WagerStorage) GetPendingWagers(
	ctx context.Context,
	walletAddress string,
	pageOpts *pymwymi.PageOpts,
) ([]pymwymi.PersistedWager, *pymwymi.Error) {
	result := []pymwymi.PersistedWager{}
	filter := bson.D{
		bson.E{Key: "status", Value: pymwymi.StatePending},
		bson.E{Key: "participants.walletAddress", Value: walletAddress},
	}
	options := setPageOptions(options.Find(), pageOpts)
	cursor, err := s.c.Find(ctx, filter, options)
	if err != nil {
		return result, pymwymi.Errorf(pymwymi.ErrInternal, "failed to get pending wagers: %s", err.Error())
	}
	for cursor.Next(ctx) {
		var w pymwymi.PersistedWager
		if err := cursor.Decode(&w); err != nil {
			log.Printf("Failed to decode: %v\n", err)
			continue
		}
		result = append(result, w)
	}
	return result, nil
}

// if you pass 0 for page and limit it will return all
func setPageOptions(opts *options.FindOptionsBuilder, p *pymwymi.PageOpts) *options.FindOptionsBuilder {
	if p.Page != 0 && p.Limit != 0 {
		opts.SetSkip((p.Page - 1) * p.Limit)
	}
	if p.Limit != 0 {
		opts.SetLimit(p.Limit)
	}
	return opts
}
