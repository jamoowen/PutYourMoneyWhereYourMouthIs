package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/challenge"
)

type Server struct {
	router *chi.Mux
	cs     *challenge.ChallengeService
	bs     *blockchain.BlockchainService
	// Db, config can be added here
}

// need auth middleware - get user to sign transaction

func CreateNewServer() *Server {
	s := &Server{}
	s.router = chi.NewRouter()
	// set content type
	// rate limit?
	// auth
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.AllowContentType("application/json"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Mount("/challenge", s.getChallengeRoutes())

	return s
}
