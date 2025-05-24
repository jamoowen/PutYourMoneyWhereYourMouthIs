package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
)

func (s *Server) mountAuthRoutes() chi.Router {
	r := chi.NewRouter()
	return r
}

func (s *Server) authenticate(w http.ResponseWriter, r *http.Request) {
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
	isSigValid, err := blockchain.AuthenticateSignature(authDTO.WalletAddress, authDTO.Sig)
	if err != nil {
		handleHttpError(w, err, http.StatusInternalServerError)
		return
	}
	if !isSigValid {
		handleHttpError(w, fmt.Errorf("signature is invalid"), http.StatusUnauthorized)
		return
	}
	jwt, err := s.authService.CreateUserJwt(pymwymi.User{WalletAddress: authDTO.WalletAddress})
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
