package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/auth"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/challenge"
)

type Server struct {
	router *chi.Mux
	// Db, config can be added here
}

// need auth middleware - get user to sign transaction
func CreateNewServer(cS *challenge.ChallengeService, bS *blockchain.BlockchainService, aS *auth.AuthService) *Server {
	s := &Server{
		router: chi.NewRouter(),
	}
	// set content type
	// rate limit?
	// auth
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.AllowContentType("application/json"))

	requireAuth := authMiddleware(aS)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.router.Use(middleware.Timeout(60 * time.Second))

	aR := authRoutes{
		authService: aS,
	}
	s.router.Mount("/auth", aR.mountAuthRoutes())

	cR := challengeRoutes{
		challengeService: cS,
	}
	s.router.Mount("/challenge", cR.mountChallengeRoutes())

	return s
}
