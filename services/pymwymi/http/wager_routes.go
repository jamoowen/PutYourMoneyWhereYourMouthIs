package http

import (
	"encoding/json"
	"fmt"
	"log"
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

type NewWagerPayload struct {
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
	r.With(s.authMiddleware).Post("/create", s.handleCreateWager)
	r.With(s.authMiddleware).Patch("/vote", s.handleVote)
	r.With(s.authMiddleware).Patch("/claim", s.handleVote)
	r.With(s.authMiddleware).Patch("/accept", s.handleAccept)

	s.router.Mount(prefix, r)
}

var validStatuses = []pymwymi.WagerStatus{
	pymwymi.StateCreated,   // not all users have accepted yet...
	pymwymi.StatePending,   // all users have accepted
	pymwymi.StateCancelled, // all users have voted cancel
	pymwymi.StateCompleted, // winner is able to claim
	pymwymi.StateClaimed,   // winner has claimed
}

// must path status as a query param eg /wager/list?status=1
func (s *Server) handleGetWagers(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	statusInt, convErr := strconv.Atoi(statusStr)
	if convErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", convErr.Error()), fmt.Sprintf("invalid status: %s", statusStr))
		return
	}

	isStatusValid := slices.Contains(validStatuses, pymwymi.WagerStatus(statusInt))
	if !isStatusValid {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, ""), fmt.Sprintf("invalid status: %s", statusStr))
		return
	}

	// these are not strictly necessary
	creator := r.URL.Query().Get("creator")
	winner := r.URL.Query().Get("winner")

	ctx := r.Context()
	pageOpts := pymwymi.GetPageOptsFromCtx(ctx)
	walletAddress := pymwymi.GetUserFromCtx(ctx).WalletAddress

	wagers, err := s.wagerService.GetWagersForUser(ctx, creator, pageOpts, pymwymi.WagerStatus(statusInt), walletAddress, winner)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to create wager")
	}
	more := true
	if int64(len(wagers)) < pageOpts.Limit {
		more = false
	}
	pagination := &Pagination{
		More:        more,
		CurrentPage: pageOpts.Page,
	}

	w.WriteHeader(http.StatusOK)
	NewPYMWYMIResponse(ctx, wagers, pagination, w)
}

func (s *Server) handleCreateWager(w http.ResponseWriter, r *http.Request) {
	// need to pass from user, participants, name, description, currency, amount, transactionHash,
	// 10 Kib
	var newWagerPayload NewWagerPayload

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	if decodeErr := json.NewDecoder(r.Body).Decode(&newWagerPayload); decodeErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", decodeErr.Error()), "bad payload")
		return
	}

	var maxStakeAmount int64 = s.blockchainService.ToWeiUSDC(1_000_000) // 1 mil
	var minStakeAmount int64 = s.blockchainService.ToWeiUSDC(1)         // 5 dollars

	stake, intError := strconv.ParseInt(newWagerPayload.Stake, 10, 64)
	if intError != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", intError.Error()), "stake must be an integer in the smallest units of USDC")
		return
	}
	if stake < minStakeAmount || stake > maxStakeAmount {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "invalid stake: %d", stake), "stake must be between $1 and $1 mil")
		return
	}
	// light validation
	// need to manually check that all the participants addresses are valid
	// need to verify the stake is the correct amount
	err := ValidateAll(
		NewStringValidator("transactionHash", newWagerPayload.TransactionHash, CheckMaxChars(66), CheckMinChars(66)),
		NewStringValidator("name", newWagerPayload.Name, CheckMaxChars(50), CheckMinChars(3)),
		NewStringValidator("category", newWagerPayload.Category, CheckMaxChars(50), CheckMinChars(3)),
		NewStringValidator("description", newWagerPayload.Description, CheckMaxChars(500)),
		NewStringValidator("location", newWagerPayload.Location, CheckMaxChars(500)),
		NewStringValidator("currency", newWagerPayload.Currency, CheckMaxChars(500), CheckMinChars(3)),
	)
	if err != nil {
		log.Printf("err: %v\npayload: %v", err, newWagerPayload)
		handlePYMWYMIError(w, err, fmt.Sprintf("bad payload: %s", err.Error()))
		return
	}

	ctx := r.Context()

	wager, err := s.wagerService.CreateWager(ctx,
		newWagerPayload.Category,
		newWagerPayload.Currency,
		newWagerPayload.Description,
		newWagerPayload.Location,
		newWagerPayload.Name,
		newWagerPayload.ParticipantsAddresses,
		stake,
		newWagerPayload.TransactionHash,
	)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to create wager")
		return
	}

	w.WriteHeader(http.StatusCreated)
	NewPYMWYMIResponse(ctx, wager, nil, w)
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
