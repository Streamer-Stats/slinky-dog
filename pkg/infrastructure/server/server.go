package server

import (
	"leagueapi.com.br/rest/application/controllers/handleController"
	"leagueapi.com.br/rest/restapi/operations"
)



// Server entrypoint for aplication
type Server struct {
	handleController *handlecontrollers.HandleController
	api *operations.LeagueAPIAPI
}

// Controllers my server run
func (s *Server) Controllers() {
	s.handleController.Handle(s.api)
}

// BindAPI bind a api from swagger
func (s *Server) BindAPI(ap *operations.LeagueAPIAPI) *Server {
	s.api = ap
	return s
}

// NewServer Server IoC
func NewServer(_handleController *handlecontrollers.HandleController) *Server {
	return &Server{
		handleController: _handleController,
	}
}
