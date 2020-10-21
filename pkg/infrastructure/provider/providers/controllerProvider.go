package providers

import (
	hc"leagueapi.com.br/rest/application/controllers/handleController"
	"leagueapi.com.br/rest/application/controllers"
	"go.uber.org/dig"

)

// ControllerProvider is a provider for services
type ControllerProvider struct {
}

// Provide is a helper Ioc
func (provider *ControllerProvider) Provide(container *dig.Container) {
	container.Provide(controllers.NewPlayerController)
	container.Provide(controllers.NewWebSocketController)
	container.Provide(hc.NewHandleController)

}

// NewControllerProvider IoC
func NewControllerProvider() *ControllerProvider {
	return &ControllerProvider{}
}
