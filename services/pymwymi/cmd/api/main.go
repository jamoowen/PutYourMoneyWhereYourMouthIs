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
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/user"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/wager"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env", err)
	}
}

func main() {
	fmt.Println("starting up wager api")
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

	wagerStorage := mongo.NewWagerStore(mongoClient, dbName)
	userStorage := mongo.NewUsersStore(mongoClient, dbName)
	wagerService := wager.NewWagerService(wagerStorage, userStorage)

	blockchainService := blockchain.NewBlockchainService()

	jwtTokenExpiration := time.Hour * 24 * 7
	authService := auth.GetAuthService(os.Getenv("JWT_SECRET"), jwtTokenExpiration)

	userService := user.NewUserService(userStorage)

	server := http.NewServer(userService, wagerService, blockchainService, authService)
	server.Start(os.Getenv("PORT"))
}

func validateEnvVars(envVars ...string) {
	for _, v := range envVars {
		if v == "" {
			log.Fatal("Missing env var")
		}
	}
}
