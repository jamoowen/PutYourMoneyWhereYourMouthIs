package http

import (
	"log"
	"net/http"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

func handleHttpError(w http.ResponseWriter, err error, code int) {
	if code >= 500 {
		log.Println(err.Error()) // Only log server-side errors
		http.Error(w, "internal server error", http.StatusInternalServerError)
	} else {
		http.Error(w, err.Error(), code)
	}
}

func handlePYMWYMIError(w http.ResponseWriter, err error) {
	switch pymwymi.GetErrorCode(err) {
	case pymwymi.ErrNotParticipant:
		http.Error(w, pymwymi.ErrorMessage(err), http.StatusBadRequest)
	case pymwymi.ErrVotingFinished:
		http.Error(w, pymwymi.ErrorMessage(err), http.StatusBadRequest)
	case pymwymi.ErrBadInput:
		http.Error(w, pymwymi.ErrorMessage(err), http.StatusBadRequest)
	case pymwymi.ErrChallengeNotFound:
		http.Error(w, pymwymi.ErrorMessage(err), http.StatusNotFound)
	default:
		log.Println(err.Error()) // Only log server-side errors
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
