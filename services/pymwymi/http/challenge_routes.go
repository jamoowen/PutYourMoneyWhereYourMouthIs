package http

import (
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
	s.router.Mount(prefix, r)
}

// must path status as a query param eg /challenge/list?status=1
func (s *Server) handleGetChallenges(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	statusInt, err := strconv.Atoi(statusStr)
	if err != nil {
		handleHttpError(w, &HttpError{Error: nil, Message: fmt.Sprintf("invalid status: %s", statusStr), Code: http.StatusBadRequest})
		return
	}
	supportedStatuses := []pymwymi.ChallengeStatus{
		pymwymi.StateCreated,
		pymwymi.StatePending,
		pymwymi.StateCancelled,
		pymwymi.StateCompleted,
		pymwymi.StateClaimed,
	}
	isSupported := slices.Contains(supportedStatuses, pymwymi.ChallengeStatus(statusInt))
	if !isSupported {
		handleHttpError(w, &HttpError{Error: nil, Message: fmt.Sprintf("invalid status: %s", statusStr), Code: http.StatusBadRequest})
		return
	}
	s.challengeService.GetChallengesForUser(r.Context(), pymwymi.ChallengeStatus(statusInt))
}
