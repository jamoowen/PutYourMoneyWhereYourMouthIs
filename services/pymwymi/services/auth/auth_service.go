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

type Service struct {
	signingKey    []byte
	durationValid time.Duration
}

func GetAuthService(signingKey string, durationValid time.Duration) *Service {
	if signingKey == "" {
		panic("signingKey is required")
	}

	return &Service{
		signingKey:    []byte(signingKey),
		durationValid: durationValid,
	}
}

type userCustomClaims struct {
	pymwymi.User
	jwt.RegisteredClaims
}

func (a *Service) CreateUserJwt(user pymwymi.User) (string, *pymwymi.Error) {
	claims := userCustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.durationValid)),
			Issuer:    "pymwymi",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwt, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", pymwymi.Errorf(pymwymi.ErrBadInput, "failed to sign jwt: %w", err)
	}
	return signedJwt, nil
}

// dont think im handling erorrs correctly not sure what happens when expired
// need to just run some tests
func (a *Service) AuthenticateUserToken(tokenString string) (pymwymi.User, *InvalidAuthError) {
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
