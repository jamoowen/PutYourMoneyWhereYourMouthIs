package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ChallengeTestSuite struct {
	suite.Suite
	server *Server
}

func (suite *ChallengeTestSuite) SetupTest() {
	s := CreateNewServer()
	s.MountHandlers()
	suite.server = s
}

func (s *ChallengeTestSuite) TestCreateChallenge() {
	req, _ := http.NewRequest("POST", "/challenge", nil)
	response := executeRequest(req, s.server)
	checkResponseCode(s.T(), http.StatusCreated, response.Code)
	require.Equal(s.T(), "Challenge created", response.Body.String())
}

func (s *ChallengeTestSuite) TestAcceptChallenge() {
	challengeId := "1"
	req, _ := http.NewRequest("PATCH", "/challenge/"+challengeId+"/accept", nil)
	response := executeRequest(req, s.server)
	checkResponseCode(s.T(), http.StatusNoContent, response.Code)
}

func (s *ChallengeTestSuite) TestVote() {
	challengeId := "1"
	req, _ := http.NewRequest("PATCH", "/challenge/"+challengeId+"/vote", nil)
	response := executeRequest(req, s.server)
	checkResponseCode(s.T(), http.StatusNoContent, response.Code)
}

func (s *ChallengeTestSuite) TestClaimChallenge() {
	challengeId := "1"
	req, _ := http.NewRequest("PATCH", "/challenge/"+challengeId+"/claim", nil)
	response := executeRequest(req, s.server)
	checkResponseCode(s.T(), http.StatusNoContent, response.Code)
}

func TestChallengeTestSuite(t *testing.T) {
	suite.Run(t, new(ChallengeTestSuite))
}

// here i need to implement the interface of my database
// func getMockDatabase() {
// 	return
// }
