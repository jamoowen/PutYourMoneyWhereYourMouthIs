package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/challenge"
)

type challengeRoutes struct {
	challengeService *challenge.ChallengeService
}

func (c *challengeRoutes) mountChallengeRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/completed", c.getCompletedChallenges)
	// r.Get("/invited", handlerFn http.HandlerFunc)
	// r.Get("/pending", handlerFn http.HandlerFunc)

	return r

	// s.router.Mount(pattern string, handler http.Handler)
}

// do i want different routes for all challenges?

func (c *challengeRoutes) getCompletedChallenges(w http.ResponseWriter, r *http.Request) {
	// completedChallenges, err := s.challengeService.getChallenges(user.walletAddress, pymwymi.StateCompleted)
}
