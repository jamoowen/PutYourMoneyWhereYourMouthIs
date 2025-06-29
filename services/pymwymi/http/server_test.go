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
// 	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/wager"
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
// type WagerTestSuite struct {
// 	suite.Suite
// 	server *Server
// }
//
// func (suite *WagerTestSuite) SetupTest() {
// 	dB := getMockDatabase()
// 	wagerDb:=
// 	cS := wager.NewWagerService(dB)
// 	bS := &blockchain.Service{}
// 	aS := &auth.Service{}
// 	uS := &user.NewUserService()
//
// 	s := NewServer(cS, bS, aS)
// 	suite.server = s
// }
//
// func (s *WagerTestSuite) TestCreateWager() {
// 	req, _ := http.NewRequest("POST", "/wager", nil)
// 	response := executeRequest(req, s.server)
// 	checkResponseCode(s.T(), http.StatusCreated, response.Code)
// 	require.Equal(s.T(), "Wager created", response.Body.String())
// }
//
// func TestWagerTestSuite(t *testing.T) {
// 	suite.Run(t, new(WagerTestSuite))
// }
//
// // CreateWager(ctx context.Context, wager *pymwymi.Wager) error
// // UpdateWager(ctx context.Context, id string, fieldsToSet []MongoField) error
// // GetWagersByStatus(walletAddress string, status pymwymi.WagerStatus, pageOpts pymwymi.PageOpts) ([]pymwymi.PersistedWager, error)
//
// func getMockDatabase() *mockDatabase {
// 	return &mockDatabase{}
// }
//
// type mockDatabase struct {
// 	wagers map[string]pymwymi.PersistedWager
// }
//
// func (s *mockDatabase) CreateWager(ctx context.Context, wager *pymwymi.Wager) error {
// 	return nil
// }
//
// func (s *mockDatabase) UpdateWager(ctx context.Context, id string, fieldsToSet []pymwymi.FieldToSet) error {
// 	return nil
// }
//
// func (s *mockDatabase) GetWagersByStatus(ctx context.Context, walletAddress string, status pymwymi.WagerStatus, pageOpts pymwymi.PageOpts) ([]pymwymi.PersistedWager, error) {
// 	return nil, nil
// }
//
// // here i need to implement the interface of my database
// // func getMockDatabase() {
// // 	return
// // }
