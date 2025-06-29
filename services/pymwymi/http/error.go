package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

type ErrorResponse struct {
	Code    pymwymi.ErrorCode `json:"code"`
	Message string            `json:"message,omitempty"`
}

func handlePYMWYMIError(w http.ResponseWriter, err *pymwymi.Error, msg string) {
	code := pymwymi.GetErrorCode(err)

	w.Header().Set("Content-Type", "application/json")

	switch code {
	case pymwymi.ErrBadInput,
		pymwymi.ErrNotParticipant,
		pymwymi.ErrNotPYMWYMIUser,
		pymwymi.ErrVotingFinished:
		log.Printf("%s: %s", msg, pymwymi.ErrorMessage(err))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:    code,
			Message: msg,
		})

	case pymwymi.ErrUserNotFound,
		pymwymi.ErrWagerNotFound:
		log.Printf("%s: %s", msg, pymwymi.ErrorMessage(err))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:    code,
			Message: msg,
		})

	default:
		log.Printf("%s: %s", msg, pymwymi.ErrorMessage(err)) // Only log internal server errors
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:    pymwymi.ErrInternal, // Optional: define this sentinel error
			Message: "internal server error",
		})
	}
}
