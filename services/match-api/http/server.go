package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	// set content type
	// rate limit?
	// auth
	//
	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)

	// Mount all handlers here
	// must be ablt
	// 1. Must be able to deposit and create a new challenge
	// 2. Challenged player must be able to accept
	// 3. Both players must be able to submit the results of the challenge
	// 4. If cancelled, players must be able to claim refund
	// 5. Winning player must be able to claim reward
	// 6. Must be able to see challenge requests
	// 7. Must be able to see past challenges & status
	// 	s.Router.Get("/", HelloWorld)
}

func (s *Server) getChallenges() {
	return
}
