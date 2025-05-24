package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

type InvalidAuthCode string

const (
	InvalidToken InvalidAuthCode = "invalid_token"
	ExpiredToken InvalidAuthCode = "expired_token"
)

type InvalidAuthError struct {
	Message string
	Code    InvalidAuthCode
}

func (e *InvalidAuthError) Error() string {
	return e.Message
}

type AuthService struct {
	signingKey    []byte
	durationValid time.Duration
}

func GetAuthService(signingKey string, durationValid time.Duration) *AuthService {
	if signingKey == "" {
		panic("signingKey is required")
	}
	return &AuthService{
		signingKey:    []byte(signingKey),
		durationValid: durationValid,
	}
}

type userCustomClaims struct {
	pymwymi.User
	jwt.RegisteredClaims
}

func (a *AuthService) CreateUserJwt(user pymwymi.User) (string, error) {
	claims := userCustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.durationValid)),
			Issuer:    "pymwymi",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.signingKey)
}

// dont think im handling erorrs correctly not sure what happens when expired
// need to just run some tests
func (a *AuthService) AuthenticateUserToken(tokenString string) (pymwymi.User, *InvalidAuthError) {
	var user pymwymi.User
	token, err := jwt.ParseWithClaims(tokenString, &userCustomClaims{}, func(token *jwt.Token) (any, error) {
		return a.signingKey, nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return user, &InvalidAuthError{Message: "token expired", Code: ExpiredToken}
	}
	if err != nil {
		return user, &InvalidAuthError{Message: "invalid token", Code: InvalidToken}
	}
	claims, ok := token.Claims.(*userCustomClaims)
	if !ok || !token.Valid {
		return user, &InvalidAuthError{Message: "Unable to read auth token", Code: InvalidToken}
	}
	return claims.User, nil
}
