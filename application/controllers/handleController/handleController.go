
package handlecontrollers

import (
	"github.com/go-openapi/runtime/middleware"
	oapi"leagueapi.com.br/rest/restapi/operations/api"
	"leagueapi.com.br/rest/application/controllers"
	"leagueapi.com.br/rest/restapi/operations"
)

// HandleController handle my controller for injection
type HandleController struct {
	PlayerController *controllers.PlayerController
}

// Handle entrypoint of controllers
func (h *HandleController) Handle(api *operations.LeagueAPIAPI) {
	api.APILiveMatchHandler = oapi.LiveMatchHandlerFunc(func(params oapi.LiveMatchParams) middleware.Responder {
		return h.PlayerController.GetDamage(params)
	})
}

// NewHandleController IoC
func NewHandleController(_playerController *controllers.PlayerController) *HandleController {
	return &HandleController{
		PlayerController: _playerController,
	}
}
