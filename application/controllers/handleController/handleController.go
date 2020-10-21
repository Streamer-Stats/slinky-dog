
package handlecontrollers

import (
	"github.com/go-openapi/runtime/middleware"
	oapi"leagueapi.com.br/rest/restapi/operations/api"
	liveapi"leagueapi.com.br/rest/restapi/operations/live"
	"leagueapi.com.br/rest/application/controllers"
	"leagueapi.com.br/rest/restapi/operations"
)

// HandleController handle my controller for injection
type HandleController struct {
	PlayerController *controllers.PlayerController
	WebSocketController *controllers.WebSocketController
}

// Handle entrypoint of controllers
func (h *HandleController) Handle(api *operations.LeagueAPIAPI) {
	api.APILiveMatchHandler = oapi.LiveMatchHandlerFunc(func(params oapi.LiveMatchParams) middleware.Responder {
		return h.PlayerController.GetDamage(params)
	})

	api.LiveLiveHandler =  liveapi.LiveHandlerFunc(func(params liveapi.LiveParams) middleware.Responder {
		return h.WebSocketController.Connect(params)
	})
}

// NewHandleController IoC
func NewHandleController(_playerController *controllers.PlayerController, _websocketController *controllers.WebSocketController) *HandleController {
	return &HandleController{
		PlayerController: _playerController,
		WebSocketController: _websocketController,
	}
}
