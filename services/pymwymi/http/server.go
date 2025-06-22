package http

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/auth"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/challenge"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/user"
)

type Server struct {
	router              *chi.Mux
	authService         *auth.Service
	userService         *user.Service
	blockchainService   *blockchain.Service
	challengeService    *challenge.Service
	authMiddleware      func(http.Handler) http.Handler
	updateChallengeBusy chan bool
	// Db, config can be added here
}

func (s *Server) Start(port string) {
	log.Printf("Listening on http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, s.router))
}

// need auth middleware - get user to sign transaction
func NewServer(uS *user.Service, cS *challenge.Service, bS *blockchain.Service, aS *auth.Service) *Server {
	s := &Server{
		router:              chi.NewRouter(),
		authService:         aS,
		userService:         uS,
		blockchainService:   bS,
		challengeService:    cS,
		authMiddleware:      authMiddleware(aS),
		updateChallengeBusy: make(chan bool, 1),
	}

	s.router.Use(httprate.LimitByIP(100, time.Minute))
	s.router.Use(middleware.RedirectSlashes)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.AllowContentType("application/json"))
	s.router.Use(middleware.Timeout(60 * time.Second))
	s.router.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// this adds a paginate key to the ctx
	// value is a pymwymi.PageOpts
	s.router.Use(paginate)

	// /auth
	s.mountAuthRoutes()
	// /challenge
	s.mountChallengeRoutes()

	return s
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://localhost:4000" {
		return true
	}
	return false
}
