package singletons

import (
	"sync"
	i"leagueapi.com.br/rest/pkg/infrastructure/provider/IoC"
)

// GetIoCSingleton get the ioc singleton
func GetIoCSingleton() *i.IoC {
	var once sync.Once
	var iocSingleton *i.IoC

	once.Do(func() {
		iocSingleton = i.NewIoC()
	})
	return iocSingleton
}
