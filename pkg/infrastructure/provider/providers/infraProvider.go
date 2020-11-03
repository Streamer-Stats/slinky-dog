package providers

import (
	"go.uber.org/dig"
	"leagueapi.com.br/rest/pkg/infrastructure/database"
	"leagueapi.com.br/rest/pkg/infrastructure/server"
	"leagueapi.com.br/rest/pkg/infrastructure/socket"
)

// InfraProvider is a provider for services
type InfraProvider struct {
}

// Provide is a helper Ioc
func (provider *InfraProvider) Provide(container *dig.Container) {
	container.Provide(socket.NewSocket)
	container.Provide(database.NewDatabase)
	container.Provide(server.NewServer)

}

// NewInfraProvider IoC
func NewInfraProvider() *InfraProvider {
	return &InfraProvider{}
}
