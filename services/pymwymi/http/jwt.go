package http

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

type InvalidJWTError struct {
	message string
}

func (e *InvalidJWTError) Error() string {
	return e.message
}

type jwtService struct {
	signingKey    string
	durationValid time.Duration
}

func newJWTService(signingKey string, durationValid time.Duration) *jwtService {
	if signingKey == "" {
		panic("signingKey is required")
	}
	return &jwtService{
		signingKey:    signingKey,
		durationValid: durationValid,
	}
}

func (j *jwtService) createUserJWT(user pymwymi.User) (string, error) {
	type MyCustomClaims struct {
		pymwymi.User
		jwt.RegisteredClaims
	}
	claims := MyCustomClaims{
		user,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.durationValid)),
			Issuer:    "pymwymi",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

// dont think im handling erorrs correctly not sure what happens when expired
func (j *jwtService) decodeJWT(tokenString string) (pymwymi.User, error) {
	var user pymwymi.User
	type MyCustomClaims struct {
		pymwymi.User
		jwt.RegisteredClaims
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.signingKey), nil
	})
	if err != nil {
		return user, fmt.Errorf("failed to parse token: %w", err)
	}
	claims, ok := token.Claims.(*MyCustomClaims)
	if claims != nil && claims.ExpiresAt.Before(time.Now()) {
		return user, &InvalidJWTError{message: "token expired"}
	}
	if !ok || !token.Valid {
		return user, &InvalidJWTError{message: "invalid token"}
	}
	return claims.User, nil
}
