package providers

import (
	"go.uber.org/dig"
	"leagueapi.com.br/rest/pkg/infrastructure/repository"
)

// RepositoryProvider is a provider for services
type RepositoryProvider struct {
}

// Provide is a helper Ioc
func (provider *RepositoryProvider) Provide(container *dig.Container) {
	container.Provide(repository.NewAuthRepository)

}

// NewRepositoryProvider IoC
func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{}
}
