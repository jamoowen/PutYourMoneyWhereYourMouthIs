package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/http"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/auth"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/challenge"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/user"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env", err)
	}
}

func main() {
	fmt.Println("starting up challenge api")
	// load env vars
	validateEnvVars(
		os.Getenv("PORT"),
		os.Getenv("MONGO_URI"),
		os.Getenv("JWT_SECRET"),
	)

	mongoClient := mongo.ConnectToMongo(os.Getenv("MONGO_URI"))
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Println("Mongo disconnect error:", err)
		}
	}()

	dbName := "pymwymi"

	challengeStorage := mongo.NewChallengeStore(mongoClient, dbName)
	userStorage := mongo.NewUsersStore(mongoClient, dbName)
	challengeService := challenge.NewChallengeService(challengeStorage, userStorage)

	blockchainService := blockchain.NewBlockchainService()

	jwtTokenExpiration := time.Hour * 24 * 7
	authService := auth.GetAuthService(os.Getenv("JWT_SECRET"), jwtTokenExpiration)

	userService := user.NewUserService(userStorage)

	server := http.NewServer(userService, challengeService, blockchainService, authService)
	server.Start(os.Getenv("PORT"))
}

func validateEnvVars(envVars ...string) {
	for _, v := range envVars {
		if v == "" {
			log.Fatal("Missing env var")
		}
	}
}
