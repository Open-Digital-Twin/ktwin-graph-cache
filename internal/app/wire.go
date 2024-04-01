package app

import "github.com/google/wire"

type AppContainer struct{}

func NewAppContainer() AppContainer {
	return AppContainer{}
}

func InitializeAppContainer() AppContainer {
	wire.Build()
	return AppContainer{}
}
