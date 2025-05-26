package http

//
// import (
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
//
// 	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
// 	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/auth"
// 	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/blockchain"
// 	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/challenge"
// 	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/user"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// )
//
// // executeRequest, creates a new ResponseRecorder
// // then executes the request by calling ServeHTTP in the router
// // after which the handler writes the response to the response recorder
// // which we can then inspect.
// func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	s.router.ServeHTTP(rr, req)
//
// 	return rr
// }
//
// // checkResponseCode is a simple utility to check the response code
// // of the response
// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }
//
// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
// type ChallengeTestSuite struct {
// 	suite.Suite
// 	server *Server
// }
//
// func (suite *ChallengeTestSuite) SetupTest() {
// 	dB := getMockDatabase()
// 	challengeDb:=
// 	cS := challenge.NewChallengeService(dB)
// 	bS := &blockchain.Service{}
// 	aS := &auth.Service{}
// 	uS := &user.NewUserService()
//
// 	s := NewServer(cS, bS, aS)
// 	suite.server = s
// }
//
// func (s *ChallengeTestSuite) TestCreateChallenge() {
// 	req, _ := http.NewRequest("POST", "/challenge", nil)
// 	response := executeRequest(req, s.server)
// 	checkResponseCode(s.T(), http.StatusCreated, response.Code)
// 	require.Equal(s.T(), "Challenge created", response.Body.String())
// }
//
// func TestChallengeTestSuite(t *testing.T) {
// 	suite.Run(t, new(ChallengeTestSuite))
// }
//
// // CreateChallenge(ctx context.Context, challenge *pymwymi.Challenge) error
// // UpdateChallenge(ctx context.Context, id string, fieldsToSet []MongoField) error
// // GetChallengesByStatus(walletAddress string, status pymwymi.ChallengeStatus, pageOpts pymwymi.PageOpts) ([]pymwymi.PersistedChallenge, error)
//
// func getMockDatabase() *mockDatabase {
// 	return &mockDatabase{}
// }
//
// type mockDatabase struct {
// 	challenges map[string]pymwymi.PersistedChallenge
// }
//
// func (s *mockDatabase) CreateChallenge(ctx context.Context, challenge *pymwymi.Challenge) error {
// 	return nil
// }
//
// func (s *mockDatabase) UpdateChallenge(ctx context.Context, id string, fieldsToSet []pymwymi.FieldToSet) error {
// 	return nil
// }
//
// func (s *mockDatabase) GetChallengesByStatus(ctx context.Context, walletAddress string, status pymwymi.ChallengeStatus, pageOpts pymwymi.PageOpts) ([]pymwymi.PersistedChallenge, error) {
// 	return nil, nil
// }
//
// // here i need to implement the interface of my database
// // func getMockDatabase() {
// // 	return
// // }
