package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) getChallengeRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/completed", getCompletedChallenges)
	// r.Get("/invited", handlerFn http.HandlerFunc)
	// r.Get("/pending", handlerFn http.HandlerFunc)

	return r

	// s.router.Mount(pattern string, handler http.Handler)
}

// do i want different routes for all challenges?

func getCompletedChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user")

	// completedChallenges, err := s.challengeService.getChallenges(user.walletAddress, pymwymi.StateCompleted)
}
