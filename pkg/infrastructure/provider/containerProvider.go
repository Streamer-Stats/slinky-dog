package provider

import (
	"log"

	"go.uber.org/dig"
	providerinterfaces "leagueapi.com.br/rest/pkg/infrastructure/provider/interface"
	"leagueapi.com.br/rest/pkg/infrastructure/provider/providers"
	"leagueapi.com.br/rest/pkg/infrastructure/server"
	"leagueapi.com.br/rest/restapi/operations"
)

// ContainerProvider controlls of IoC
type ContainerProvider struct {
	Container *dig.Container
	providers [4]providerinterfaces.IProvider
}

func buildContainerProvider() *dig.Container {
	return dig.New()
}

// Provide My Ioc
func (p *ContainerProvider) Provide() *dig.Container {
	for _, pro := range p.providers {
		pro.Provide(p.Container)
	}

	return p.Container
}

// Run EntryPoint
func (p *ContainerProvider) Run(api *operations.LeagueAPIAPI) {
	err := p.Provide().Invoke(func(server *server.Server) {
		server.BindAPI(api).Controllers()
	})

	if err != nil {
		log.Printf("%v", err)
	}

}

// NewIocProvider Create provider
func NewIocProvider() *ContainerProvider {
	return &ContainerProvider{
		Container: buildContainerProvider(),
		providers: [4]providerinterfaces.IProvider{
			providers.NewInfraProvider(),
			providers.NewRepositoryProvider(),
			providers.NewServiceProvider(),
			providers.NewControllerProvider(),
		},
	}
}
