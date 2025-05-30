package http

import (
	"log"
	"net/http"
)

func handleHttpError(w http.ResponseWriter, err error, code int) {
	if code >= 500 {
		log.Println(err.Error()) // Only log server-side errors
		http.Error(w, "internal server error", http.StatusInternalServerError)
	} else {
		http.Error(w, err.Error(), code)
	}
}
