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

type VoteDTO struct {
	ChallengeId string       `json:"challengeId"`
	Vote        pymwymi.Vote `json:"vote"`
}

type AcceptDTO struct {
	ChallengeId    string `json:"challengeId"`
	StakeSignature string `json:"stakeSignature"`
}

func (s *Server) mountChallengeRoutes() {
	prefix := "/challenge"

	r := chi.NewRouter()
	r.With(s.authMiddleware).Get("/list", s.handleGetChallenges)
	r.With(s.authMiddleware).Get("/create", s.handleCreateChallenge)
	r.With(s.authMiddleware).Get("/vote", s.handleVote)

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

// user clicks accept -> signs transaction -> we send transaction sig with challenge id
// do we need to look up the transaction?
// this is going to be a lot more complicated than i thought
// wait does it? the smart contract ensures the amount sent was correct
// so all i need to do is verify that the transaction went through?
func (s *Server) handleAccept(w http.ResponseWriter, r *http.Request) {
	var acceptDTO AcceptDTO
	err := json.NewDecoder(r.Body).Decode(&acceptDTO)
	if err != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "bad request body: %v", err.Error()))
		return
	}

	ctx := r.Context()
	user := pymwymi.GetUserFromCtx(ctx)

	challenge, err := s.challengeService.GetChallengeForParticipant(ctx, acceptDTO.ChallengeId, user.WalletAddress)
	if err != nil {
		handlePYMWYMIError(w, err)
	}

	validStatuses := []pymwymi.ChallengeStatus{pymwymi.StatePending, pymwymi.StateCreated}
	if slices.Contains(validStatuses, challenge.Status) {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrVotingFinished, "voting no longer possible"))
	}

	s.updateChallengeBusy <- true
	defer func() { <-s.updateChallengeBusy }()

	w.WriteHeader(http.StatusNoContent)
}

// vote is the wallett address of the player they say won
func (s *Server) handleVote(w http.ResponseWriter, r *http.Request) {
	var voteDTO VoteDTO
	err := json.NewDecoder(r.Body).Decode(&voteDTO)
	if err != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "bad request body: %v", err.Error()))
		return
	}
	if voteDTO.Vote.Intent != pymwymi.VoteWinner && voteDTO.Vote.Intent != pymwymi.VoteCancel {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "invalid vote intent: %v", voteDTO.Vote.Intent))
		return
	}

	ctx := r.Context()
	user := pymwymi.GetUserFromCtx(ctx)

	challenge, err := s.challengeService.GetChallengeForParticipant(ctx, voteDTO.ChallengeId, user.WalletAddress)
	if err != nil {
		handlePYMWYMIError(w, err)
	}

	validStatuses := []pymwymi.ChallengeStatus{pymwymi.StatePending, pymwymi.StateCreated}
	if slices.Contains(validStatuses, challenge.Status) {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrVotingFinished, "voting no longer possible"))
	}

	vote := pymwymi.Vote{
		HasVoted: true,
		Intent:   voteDTO.Vote.Intent,
		Winner:   voteDTO.Vote.Winner,
	}

	s.updateChallengeBusy <- true
	defer func() { <-s.updateChallengeBusy }()
	err = s.challengeService.SubmitVote(ctx, user, challenge, vote)
	if err != nil {
		handlePYMWYMIError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
