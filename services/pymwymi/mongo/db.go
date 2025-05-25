package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectToMongo(mongoURI string) *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	// Create a new client and connect to the server
	log.Println("Connecting to MongoDB...", mongoURI)
	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetMaxPoolSize(50).
		SetMinPoolSize(5)
		// SetMaxConnIdleTime(5 * time.Minute)

	client, err := mongo.Connect(clientOptions)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}
