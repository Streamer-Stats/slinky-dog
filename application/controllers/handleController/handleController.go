package handlecontrollers

import (
	"github.com/go-openapi/runtime/middleware"
	"leagueapi.com.br/rest/application/controllers"
	"leagueapi.com.br/rest/restapi/operations"
	oapi "leagueapi.com.br/rest/restapi/operations/api"
	autheapi "leagueapi.com.br/rest/restapi/operations/auth"
	liveapi "leagueapi.com.br/rest/restapi/operations/live"
)

// HandleController handle my controller for injection
type HandleController struct {
	PlayerController    *controllers.PlayerController
	WebSocketController *controllers.WebSocketController
	AuthController      *controllers.AuthController
}

// Handle entrypoint of controllers
func (h *HandleController) Handle(api *operations.LeagueAPIAPI) {
	api.APILiveMatchHandler = oapi.LiveMatchHandlerFunc(func(params oapi.LiveMatchParams) middleware.Responder {
		return h.PlayerController.GetDamage(params)
	})

	api.LiveLiveHandler = liveapi.LiveHandlerFunc(func(params liveapi.LiveParams) middleware.Responder {
		return h.WebSocketController.Connect(params)
	})

	api.AuthAuthHandler = autheapi.AuthHandlerFunc(func(params autheapi.AuthParams) middleware.Responder {
		return h.AuthController.Login(params)
	})

}

// NewHandleController IoC
func NewHandleController(_playerController *controllers.PlayerController,
	_websocketController *controllers.WebSocketController,
	_authController *controllers.AuthController) *HandleController {
	return &HandleController{
		PlayerController:    _playerController,
		WebSocketController: _websocketController,
		AuthController:      _authController,
	}
}
