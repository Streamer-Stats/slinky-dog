package providers

import (
	"go.uber.org/dig"
	"leagueapi.com.br/rest/pkg/domain/services"
)

// ServiceProvider is a provider for services
type ServiceProvider struct {
}

// Provide is a helper Ioc
func (provider *ServiceProvider) Provide(container *dig.Container) {
	container.Provide(services.NewAuthService)

}

// NewServiceProvider IoC
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
