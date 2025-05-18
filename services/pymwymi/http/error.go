package http

import (
	"log"
	"net/http"
)

func handleHttpError(w http.ResponseWriter, err error, code int) {
	switch code {
	case http.StatusBadRequest:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case http.StatusNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	case http.StatusUnauthorized:
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case http.StatusForbidden:
		http.Error(w, err.Error(), http.StatusForbidden)
	default:
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
}
