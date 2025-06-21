package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
)

const signInString = "PYMWYMI sign in"

func (s *Server) mountAuthRoutes() {
	prefix := "/auth"

	r := chi.NewRouter()
	r.Post("/", s.authenticate)
	s.router.Mount(prefix, r)
}

// accepts walletAddress and sig and responds with a jwt in cookie if successful
func (s *Server) authenticate(w http.ResponseWriter, r *http.Request) {
	type AuthDTO struct {
		WalletAddress string `json:"walletAddress"`
		Signature     string `json:"signature"`
	}
	var authDTO AuthDTO
	err := json.NewDecoder(r.Body).Decode(&authDTO)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to decode request body: %w", err), http.StatusBadRequest)
		return
	}
	isSigValid, err := blockchain.AuthenticateSignature(authDTO.WalletAddress, authDTO.Signature, signInString)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to verify signature: %w", err), http.StatusBadRequest)
		return
	}
	if !isSigValid {
		handleHttpError(w, fmt.Errorf("invalid signature"), http.StatusBadRequest)
		return
	}
	// will create a new user if they dont exist already
	err = s.userService.CreateUser(r.Context(), authDTO.WalletAddress)
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to create user: %w", err), http.StatusInternalServerError)
		return
	}
	jwt, err := s.authService.CreateUserJwt(pymwymi.User{WalletAddress: authDTO.WalletAddress})
	if err != nil {
		handleHttpError(w, fmt.Errorf("failed to create jwt: %w", err), http.StatusInternalServerError)
		return
	}

	// TODO set secure true?
	http.SetCookie(w, &http.Cookie{
		Name:     "pymwymi_auth_token",
		Value:    jwt,
		HttpOnly: true,
		Secure:   false, // only if using HTTPS (which you should in production)
		Path:     "/",
		SameSite: http.SameSiteStrictMode, // will only use cookie if same domain
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"authenticated"}`))
}
