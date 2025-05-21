package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
)

type authRoutes struct {
	jwtService *jwtService
}

func (a *authRoutes) mountAuthRoutes() chi.Router {
	r := chi.NewRouter()

	// NOTE MUST FLIP SECURE TO TRUE IN PROD
	r.Post("/", a.authenticate)
	return r
}

func (a *authRoutes) authenticate(w http.ResponseWriter, r *http.Request) {
	type AuthDTO struct {
		WalletAddress string `json:"walletAddress"`
		Sig           string `json:"sig"`
	}
	var authDTO AuthDTO
	err := json.NewDecoder(r.Body).Decode(&authDTO)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to parse body"), http.StatusBadRequest)
		return
	}
	if authDTO.WalletAddress == "" || authDTO.Sig == "" {
		handleHttpError(w, fmt.Errorf("walletAddress and sig are required"), http.StatusBadRequest)
		return
	}
	valid, err := blockchain.AuthenticateSignature(authDTO.WalletAddress, authDTO.Sig)
	if err != nil {
		handleHttpError(w, err, http.StatusInternalServerError)
		return
	}
	if !valid {
		handleHttpError(w, fmt.Errorf("invalid signature"), http.StatusUnauthorized)
		return
	}
	jwt, err := a.jwtService.createUserJWT(pymwymi.User{Name: "Unknown", WalletAddress: authDTO.WalletAddress})
	if errors.As(err, &InvalidJWTError{}) {
		handleHttpError(w, err, http.StatusUnauthorized)
		return
	}
	if err != nil {
		handleHttpError(w, err, http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "pymwymi_auth_token",
		Value:    jwt,
		HttpOnly: true,
		Secure:   false, // only if using HTTPS (which you should in production)
		Path:     "/",
		SameSite: http.SameSiteNoneMode, // this a little dodgy might want to changee
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"authenticated"}`))
}
