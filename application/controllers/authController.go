package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"leagueapi.com.br/rest/pkg/domain/services"
	oapi "leagueapi.com.br/rest/restapi/operations/auth"
)

// AuthController controller of player acts
type AuthController struct {
	service *services.AuthService
}

// Login get damage from player
func (pc *AuthController) Login(params oapi.AuthParams) middleware.Responder {

	userWithToken := pc.service.Login(params.Body)
	return oapi.NewAuthOK().WithPayload(userWithToken)

}

// NewAuthController Constructor AuthController
func NewAuthController(_service *services.AuthService) *AuthController {
	return &AuthController{
		service: _service,
	}
}
