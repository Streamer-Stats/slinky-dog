package providers



import (
	"leagueapi.com.br/rest/pkg/infrastructure/server"
	"go.uber.org/dig"

)

// InfraProvider is a provider for services
type InfraProvider struct {
}

// Provide is a helper Ioc
func (provider *InfraProvider) Provide(container *dig.Container) {
	container.Provide(server.NewServer)

}

// NewInfraProvider IoC
func NewInfraProvider() *InfraProvider {
	return &InfraProvider{}
}
