package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
)

type AuthResponseCode string

const (
	signInString                          = "PYMWYMI_sign_in"
	authResponseSignedIn AuthResponseCode = "SIGNED_IN"
	authResponseSignedUp AuthResponseCode = "SIGNED_UP"
)

type AuthDTO struct {
	WalletAddress string `json:"walletAddress"`
	Signature     string `json:"signature"`
}

type ChangeNameDTO struct {
	Name string `json:"name"`
}

type AuthResponse struct {
	AuthCode AuthResponseCode `json:"authCode"`
}

func (s *Server) mountAuthRoutes() {
	prefix := "/auth"

	r := chi.NewRouter()
	r.Post("/", s.authenticate)
	r.Patch("/", s.updateName)
	s.router.Mount(prefix, r)
}

// accepts walletAddress and sig and responds with a jwt in cookie if successful
func (s *Server) authenticate(w http.ResponseWriter, r *http.Request) {
	var authResponseCode AuthResponseCode
	var authDTO AuthDTO
	decodeErr := json.NewDecoder(r.Body).Decode(&authDTO)
	if decodeErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", decodeErr.Error()), "bad payload")
		return
	}

	isSigValid, err := blockchain.AuthenticateSignature(authDTO.WalletAddress, authDTO.Signature, signInString)
	if err != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrInternal, "%s", err.Error()), "failed to verify signature")
		return
	}
	if !isSigValid {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, ""), "bad payload")
		return
	}

	// will create a new user if they dont exist already

	ctx := r.Context()

	var user pymwymi.User
	existingUser, err := s.userService.GetUser(ctx, authDTO.WalletAddress)
	if err != nil && err.Code != pymwymi.ErrUserNotFound {
		handlePYMWYMIError(w, err, "failed to authenticate")
		return
	}
	if err != nil && err.Code == pymwymi.ErrUserNotFound {
		newUser, err := s.userService.CreateUser(ctx, authDTO.WalletAddress)
		if err != nil {
			handlePYMWYMIError(w, err, "failed to authenticate")
			return
		}
		user = newUser
		authResponseCode = authResponseSignedUp
	}
	if err == nil {
		user.Name = existingUser.Name
		user.WalletAddress = existingUser.WalletAddress
		authResponseCode = authResponseSignedIn
	}

	jwt, err := s.authService.CreateUserJwt(user)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to authenticate")
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
	json.NewEncoder(w).Encode(AuthResponse{AuthCode: authResponseCode})
}

func (s *Server) updateName(w http.ResponseWriter, r *http.Request) {
	var authResponseCode AuthResponseCode
	var changeNameDTO ChangeNameDTO
	decodeErr := json.NewDecoder(r.Body).Decode(&changeNameDTO)
	if decodeErr != nil {
		handlePYMWYMIError(w, pymwymi.Errorf(pymwymi.ErrBadInput, "%s", decodeErr.Error()), "bad payload")
		return
	}

	ctx := r.Context()
	user := pymwymi.GetUserFromCtx(ctx)

	updatedUser, err := s.userService.UpdateName(ctx, changeNameDTO.Name, user.WalletAddress)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to update name")
		return
	}

	jwt, err := s.authService.CreateUserJwt(updatedUser)
	if err != nil {
		handlePYMWYMIError(w, err, "failed to authenticate")
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
	json.NewEncoder(w).Encode(AuthResponse{AuthCode: authResponseCode})
}
