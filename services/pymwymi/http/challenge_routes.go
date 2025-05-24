package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) mountChallengeRoutes() chi.Router {
	r := chi.NewRouter()

	r.With(s.authMiddleware).Get("/", s.getCompletedChallenges)
	// r.Get("/invited", handlerFn http.HandlerFunc)
	// r.Get("/pending", handlerFn http.HandlerFunc)

	return r
}

// do i want different routes for all challenges?

func (s *Server) getCompletedChallenges(w http.ResponseWriter, r *http.Request) {
	// get user from ctx...
	// get status from query param
	status := r.URL.Query().Get("status")

	completedChallenges, err := s.challengeService.getChallenges(user.walletAddress, pymwymi.StateCompleted)
}
