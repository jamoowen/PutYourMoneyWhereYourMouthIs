package mongo

import "github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"

type ChallengeStore interface {
	// need some sort of mongo connection here
	getChallengesForUser(walletAddress string, status pymwymi.ChallengeStatus, page, limit int)
}

type Storage struct{}

func (s *Storage) getChallengesForUser(walletAddress string, status pymwymi.ChallengeStatus, page, limit int) {
	// return challenges
}
