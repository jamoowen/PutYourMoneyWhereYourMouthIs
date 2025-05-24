package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/mongo"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/auth"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env", err)
	}
}

func main() {
	port := 4010
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// load env vars
	validateEnvVars(
		os.Getenv("MONGO_URI"),
		os.Getenv("JWT_SECRET"),
	)

	mongoClient := mongo.ConnectToMongo(os.Getenv("MONGO_URI"))

	jwtTokenExpiration := time.Hour * 24 * 7
	authService := auth.GetAuthService(os.Getenv("JWT_SECRET"), jwtTokenExpiration)

	log.Printf("Listening on http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

func validateEnvVars(envVars ...string) {
	for _, v := range envVars {
		if v == "" {
			log.Fatal("Missing env var")
		}
	}
}
