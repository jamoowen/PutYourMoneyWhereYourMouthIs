package auth

import (
	"testing"
	"time"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AuthTestSuite struct {
	suite.Suite
	authService *Service
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *AuthTestSuite) SetupTest() {
	s.authService = GetAuthService("secret", time.Minute*1)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *AuthTestSuite) TestCreateUserJwt() {
	user := pymwymi.User{
		Name:          "james",
		WalletAddress: "0x123",
	}
	_, err := s.authService.CreateUserJwt(user)
	assert.Nil(s.T(), err)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *AuthTestSuite) TestAuthenticateUserToken() {
	user := pymwymi.User{
		Name:          "james",
		WalletAddress: "0x123",
	}
	jwt, err := s.authService.CreateUserJwt(user)
	assert.Nil(s.T(), err)
	authenticatedUser, authErr := s.authService.AuthenticateUserToken(jwt)
	assert.Nil(s.T(), authErr)
	assert.Equal(s.T(), user, authenticatedUser)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *AuthTestSuite) TestBadToken() {
	_, err := s.authService.AuthenticateUserToken("nonsense")
	assert.Equal(s.T(), InvalidToken, err.Code)
}

func (s *AuthTestSuite) TestExpiredToken() {
	s.authService.durationValid = time.Second * 1
	user := pymwymi.User{
		Name:          "james",
		WalletAddress: "0x123",
	}
	jwt, err := s.authService.CreateUserJwt(user)
	assert.Nil(s.T(), err)
	time.Sleep(time.Second * 2)
	_, authErr := s.authService.AuthenticateUserToken(jwt)
	s.T().Log(authErr)
	assert.Equal(s.T(), ExpiredToken, authErr.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
