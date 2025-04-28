package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	// set content type
	// rate limit?
	// auth
	//
	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	// all of these routes need auth middleware i think
	s.Router.Use(middleware.Logger)

	s.Router.Get("/challenges", s.getChallenges)
	s.Router.Get("/challenges/created", s.getCreatedChallenges)
	s.Router.Get("/challenges/disputed", s.getDisputedChallenges)
	s.Router.Get("/challenges/claimable", s.getClaimableChallenges)
	s.Router.Get("/challenges/completed", s.getCompletedChallenges)

	s.Router.Post("/challenge", s.createChallenge)
	s.Router.Patch("/challenge/{id}", s.acceptChallenge)
	s.Router.Patch("/challenge/{id}/vote", s.vote)
	s.Router.Patch("/challenge/{id}/claim", s.claim)

	// Mount all handlers here
	// must be ablt
	// 1. Must be able to deposit and create a new challenge
	// 2. Challenged player must be able to accept
	// 3. Both players must be able to submit the results of the challenge
	// 4. If cancelled, players must be able to claim refund
	// 5. Winning player must be able to claim reward
	// 6. Must be able to see challenge requests
	// 7. Must be able to see past challenges & status
	// 	s.Router.Get("/", HelloWorld)
}

// returns all of the challenges a user has been sent
func (s *Server) getChallenges(w http.ResponseWriter, r *http.Request) {
	// get the
}

// returns all of the challenges a user has created
func (s *Server) getCreatedChallenges(w http.ResponseWriter, r *http.Request) {
	// get the
}

// get the disputed challenges
// we should give users the ability to change their vote if its disupted
func (s *Server) getDisputedChallenges(w http.ResponseWriter, r *http.Request) {
	// get the
}

// get the challenges a user can claim
func (s *Server) getClaimableChallenges(w http.ResponseWriter, r *http.Request) {
	// get the
}

// get past challanges a user has won or lost and that have been claimed
func (s *Server) getCompletedChallenges(w http.ResponseWriter, r *http.Request) {
	// get the
}

func (s *Server) createChallenge(w http.ResponseWriter, r *http.Request) {
	// get the
}

func (s *Server) acceptChallenge(w http.ResponseWriter, r *http.Request) {
	// get the
}

// vote for a winner, loser or cancel
func (s *Server) vote(w http.ResponseWriter, r *http.Request) {
	// get the
}

// how are we going to claim???
// we can either create a withdrawal slip
// or we can make the contract work in a way
// that allows a user to withdraw if the status
// matches the status
// if cancel then bothe users should be able to claim what they initially staked
func (s *Server) claim(w http.ResponseWriter, r *http.Request) {
	// get the
}
