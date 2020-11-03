package providers

import (
	"go.uber.org/dig"
	"leagueapi.com.br/rest/application/controllers"
	hc "leagueapi.com.br/rest/application/controllers/handleController"
)

// ControllerProvider is a provider for services
type ControllerProvider struct {
}

// Provide is a helper Ioc
func (provider *ControllerProvider) Provide(container *dig.Container) {
	container.Provide(controllers.NewPlayerController)
	container.Provide(controllers.NewWebSocketController)
	container.Provide(controllers.NewAuthController)
	container.Provide(hc.NewHandleController)

}

// NewControllerProvider IoC
func NewControllerProvider() *ControllerProvider {
	return &ControllerProvider{}
}
