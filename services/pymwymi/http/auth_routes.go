package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
)

const SIGN_IN_STRING = "sign-in"

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
		Sig           string `json:"sig"`
	}
	var authDTO AuthDTO
	err := json.NewDecoder(r.Body).Decode(&authDTO)
	if err != nil {
		handleHttpError(w, &HttpError{Error: err, Message: "failed to decode request body", Code: http.StatusBadRequest})
		return
	}
	isSigValid, err := blockchain.AuthenticateSignature(authDTO.WalletAddress, authDTO.Sig, SIGN_IN_STRING)
	if err != nil {
		handleHttpError(w, &HttpError{Error: err, Message: "failed to authenticate signature", Code: http.StatusBadRequest})
		return
	}
	if !isSigValid {
		handleHttpError(w, &HttpError{Error: nil, Message: "invalid signature", Code: http.StatusUnauthorized})
		return
	}
	// will create a new user if they dont exist already
	err = s.userService.CreateUser(r.Context(), authDTO.WalletAddress)
	if err != nil {
		handleHttpError(w, &HttpError{Error: err, Message: "failed to create user", Code: http.StatusInternalServerError})
		return
	}
	jwt, err := s.authService.CreateUserJwt(pymwymi.User{WalletAddress: authDTO.WalletAddress})
	if err != nil {
		handleHttpError(w, &HttpError{Error: err, Message: "failed to create jwt", Code: http.StatusInternalServerError})
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
