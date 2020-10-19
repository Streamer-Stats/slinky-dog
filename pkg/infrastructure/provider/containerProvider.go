package provider

import (
	"leagueapi.com.br/rest/pkg/infrastructure/provider/providers"
	"leagueapi.com.br/rest/pkg/infrastructure/server"
	"leagueapi.com.br/rest/restapi/operations"
	"leagueapi.com.br/rest/pkg/infrastructure/provider/interface"
	"go.uber.org/dig"
	"log"

)

// ContainerProvider controlls of IoC
type ContainerProvider struct {
	Container *dig.Container
	providers [2]providerinterfaces.IProvider
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
		providers: [2]providerinterfaces.IProvider{
			providers.NewInfraProvider(),
			providers.NewControllerProvider(),
		},
	}
}
