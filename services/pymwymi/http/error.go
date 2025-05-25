package http

import (
	"log"
	"net/http"
)

type HttpError struct {
	Error   error
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handleHttpError(w http.ResponseWriter, httpError *HttpError) {
	if httpError.Code >= 500 {
		log.Println(httpError.Error) // Only log server-side errors
		http.Error(w, "internal server error", http.StatusInternalServerError)
	} else {
		http.Error(w, httpError.Message, httpError.Code)
	}
}
