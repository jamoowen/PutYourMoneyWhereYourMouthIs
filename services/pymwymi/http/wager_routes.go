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
	WagerId string       `json:"wagerId"`
	Vote    pymwymi.Vote `json:"vote"`
}

type AcceptDTO struct {
	WagerId        string `json:"wagerId"`
	StakeSignature string `json:"stakeSignature"`
}

type NewWagerDto struct {
	TransactionHash       string   `json:"transactionHash"`
	Name                  string   `json:"name"`
	Category              string   `json:"category"`
	Description           string   `json:"description"`
	Location              string   `json:"location"`
	Stake                 string   `json:"stake"`
	Currency              string   `json:"currency"`
	ParticipantsAddresses []string `json:"participantsAddresses"`
}

func (s *Server) mountWagerRoutes() {
	prefix := "/wager"

	r := chi.NewRouter()
	r.With(s.authMiddleware).Get("/list", s.handleGetWagers)
	r.With(s.authMiddleware).Get("/create", s.handleCreateWager)
	r.With(s.authMiddleware).Get("/vote", s.handleVote)

	s.router.Mount(prefix, r)
}

var supportedStatuses = []pymwymi.WagerStatus{
	pymwymi.StateCreated,
	pymwymi.StatePending,
	pymwymi.StateCancelled,
	pymwymi.StateCompleted,
	pymwymi.StateClaimed,
}

// must path status as a query param eg /wager/list?status=1
func (s *Server) handleGetWagers(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	statusInt, err := strconv.Atoi(statusStr)
	if err != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", err.Error()), fmt.Sprintf("invalid status: %s", statusStr))
		return
	}

	isSupported := slices.Contains(supportedStatuses, pymwymi.WagerStatus(statusInt))
	if !isSupported {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, ""), fmt.Sprintf("invalid status: %s", statusStr))
		return
	}
	s.wagerService.GetWagersForUser(r.Context(), pymwymi.WagerStatus(statusInt))
}

func (s *Server) handleCreateWager(w http.ResponseWriter, r *http.Request) {
	// need to pass from user, participants, name, description, currency, amount, transactionHash,
	// 10 Kib
	var c NewWagerDto

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	if decodeErr := json.NewDecoder(r.Body).Decode(&c); decodeErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", decodeErr.Error()), "bad payload")
		return
	}

	// light validation
	// need to manually check that all the participants addresses are valid
	// need to verify the stake is the correct amount
	err := ValidateAll(
		NewStringValidator("transactionHash", c.TransactionHash, CheckMaxChars(66), CheckMinChars(66)),
		NewStringValidator("name", c.Name, CheckMaxChars(50), CheckMinChars(5)),
		NewStringValidator("category", c.Category, CheckMaxChars(50), CheckMinChars(5)),
		NewStringValidator("description", c.Description, CheckMaxChars(500), CheckMinChars(5)),
		NewStringValidator("location", c.Location, CheckMaxChars(500), CheckMinChars(5)),
		NewStringValidator("stake", c.Stake, CheckMaxChars(50), CheckMinChars(5)),
		NewStringValidator("currency", c.Currency, CheckMaxChars(500), CheckMinChars(5)),
	)
	if err != nil {
		handlePYMWYMIError(w, err, fmt.Sprintf("bad payload: %s", err.Error()))
		return
	}

	wager, err := s.wagerService.CreateWager(r.Context(),
		c.TransactionHash,
		c.Name,
		c.Category,
		c.Description,
		c.Location,
		c.Stake,
		c.Currency,
		c.ParticipantsAddresses,
	)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to create wager")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wager)
}

// user clicks accept -> signs transaction -> we send transaction sig with wager id
// do we need to look up the transaction?
// this is going to be a lot more complicated than i thought
// wait does it? the smart contract ensures the amount sent was correct
// so all i need to do is verify that the transaction went through?
func (s *Server) handleAccept(w http.ResponseWriter, r *http.Request) {
	var acceptDTO AcceptDTO
	decodeErr := json.NewDecoder(r.Body).Decode(&acceptDTO)
	if decodeErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", decodeErr.Error()), "bad payload")
		return
	}

	ctx := r.Context()
	user := pymwymi.GetUserFromCtx(ctx)

	wager, err := s.wagerService.GetWagerForParticipant(ctx, acceptDTO.WagerId, user.WalletAddress)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to get wager")
	}

	validStatuses := []pymwymi.WagerStatus{pymwymi.StatePending, pymwymi.StateCreated}
	if slices.Contains(validStatuses, wager.Status) {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrVotingFinished, "%s", err.Error()), "voting no longer possible")
	}

	s.updateWagerBusy <- true
	defer func() { <-s.updateWagerBusy }()

	w.WriteHeader(http.StatusNoContent)
}

// vote is the wallett address of the player they say won
func (s *Server) handleVote(w http.ResponseWriter, r *http.Request) {
	var voteDTO VoteDTO
	decodeErr := json.NewDecoder(r.Body).Decode(&voteDTO)
	if decodeErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", decodeErr.Error()), "bad payload")
		return
	}
	if voteDTO.Vote.Intent != pymwymi.VoteWinner && voteDTO.Vote.Intent != pymwymi.VoteCancel {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, ""), fmt.Sprintf("invalid vote intent: %v", voteDTO.Vote.Intent))
		return
	}

	ctx := r.Context()
	user := pymwymi.GetUserFromCtx(ctx)

	wager, err := s.wagerService.GetWagerForParticipant(ctx, voteDTO.WagerId, user.WalletAddress)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to get wager")
	}

	validStatuses := []pymwymi.WagerStatus{pymwymi.StatePending, pymwymi.StateCreated}
	if slices.Contains(validStatuses, wager.Status) {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrVotingFinished, ""), "voting no longer possible")
	}

	vote := pymwymi.Vote{
		HasVoted: true,
		Intent:   voteDTO.Vote.Intent,
		Winner:   voteDTO.Vote.Winner,
	}

	s.updateWagerBusy <- true
	defer func() { <-s.updateWagerBusy }()
	err = s.wagerService.SubmitVote(ctx, user, wager, vote)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to vote")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
