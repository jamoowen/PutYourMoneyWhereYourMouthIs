package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
)

func (s *Server) getAuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", authorize)
	// r.Get("/invited", handlerFn http.HandlerFunc)
	// r.Get("/pending", handlerFn http.HandlerFunc)

	return r

	// s.router.Mount(pattern string, handler http.Handler)
}

// do i want different routes for all challenges?

func authorize(w http.ResponseWriter, r *http.Request) {
	// 1. parse the body
	// 2. extract walletAddress & sig
	// 3. encode jwt with walletAddress
	// 4. return jwt
	type AuthDTO struct {
		WalletAddress string `json:"walletAddress"`
		Sig           string `json:"sig"`
	}
	var authDTO AuthDTO
	err := json.NewDecoder(r.Body).Decode(&authDTO)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to parse body"), http.StatusBadRequest)
		return
	}
	if authDTO.WalletAddress == "" || authDTO.Sig == "" {
		handleHttpError(w, fmt.Errorf("walletAddress and sig are required"), http.StatusBadRequest)
		return
	}
	valid, err := blockchain.AuthenticateSignature(authDTO.WalletAddress, authDTO.Sig)
	if err != nil {
		handleHttpError(w, err, http.StatusInternalServerError)
		return
	}
	if !valid {
		handleHttpError(w, fmt.Errorf("invalid signature"), http.StatusUnauthorized)
		return
	}
	// create jwt
	// completedChallenges, err := s.challengeService.getChallenges(user.walletAddress, pymwymi.StateCompleted)
}
