package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

func (s *Server) mountChallengeRoutes() {
	prefix := "/challenge"

	r := chi.NewRouter()
	r.With(s.authMiddleware).Get("/list", s.handleGetChallenges)
	r.With(s.authMiddleware).Get("/create", s.handleCreateChallenge)

	s.router.Mount(prefix, r)
}

var supportedStatuses = []pymwymi.ChallengeStatus{
	pymwymi.StateCreated,
	pymwymi.StatePending,
	pymwymi.StateCancelled,
	pymwymi.StateCompleted,
	pymwymi.StateClaimed,
}

// must path status as a query param eg /challenge/list?status=1
func (s *Server) handleGetChallenges(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	statusInt, err := strconv.Atoi(statusStr)
	if err != nil {
		handleHttpError(w, fmt.Errorf("invalid status: %s", statusStr), http.StatusBadRequest)
		return
	}

	isSupported := slices.Contains(supportedStatuses, pymwymi.ChallengeStatus(statusInt))
	if !isSupported {
		handleHttpError(w, fmt.Errorf("invalid status: %s", statusStr), http.StatusBadRequest)
		return
	}
	s.challengeService.GetChallengesForUser(r.Context(), pymwymi.ChallengeStatus(statusInt))
}

func (s *Server) handleCreateChallenge(w http.ResponseWriter, r *http.Request) {
	// need to pass from user, participants, name, description, currency, amount, transactionHash,
	// 10 Kib
	var c pymwymi.NewChallengeDto
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		handleHttpError(w, fmt.Errorf("failed to decode request body: %w", err), http.StatusBadRequest)
		return
	}

	// light validation
	// need to manually check that all the participants addresses are valid
	// need to verify the stake is the correct amount
	err := ValidateAll(
		NewStringValidator("transactionHash", c.TransactionHash, CheckMaxChars(66), CheckMinChars(66)),
		NewStringValidator("creator", c.Creator, CheckMaxChars(66), CheckMinChars(66)),
		NewStringValidator("name", c.Name, CheckMaxChars(50), CheckMinChars(5)),
		NewStringValidator("category", c.Category, CheckMaxChars(50), CheckMinChars(5)),
		NewStringValidator("description", c.Description, CheckMaxChars(500), CheckMinChars(5)),
		NewStringValidator("location", c.Location, CheckMaxChars(500), CheckMinChars(5)),
		NewStringValidator("stake", c.Stake, CheckMaxChars(50), CheckMinChars(5)),
		NewStringValidator("currency", c.Currency, CheckMaxChars(500), CheckMinChars(5)),
	)
	if err != nil {
		handleHttpError(w, fmt.Errorf("bad input: %w", err), http.StatusBadRequest)
		return
	}

	challenge, err := s.challengeService.CreateChallenge(r.Context(), c)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to create challenge: %w", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(challenge)
}

// vote is the wallett address of the player they say won
func (s *Server) handleVote(w http.ResponseWriter, r *http.Request) {
	var voteDTO pymwymi.VoteDTO
	err := json.NewDecoder(r.Body).Decode(&voteDTO)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to decode request body: %w", err), http.StatusBadRequest)
		return
	}
	checkIsInList := IsInList()
	err = checkIsInList(string(voteDTO.Vote.Status), string(pymwymi.VoteCancel), string(pymwymi.VoteWinner))
	if err != nil {
		handleHttpError(w, fmt.Errorf("invalid vote: %w", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
