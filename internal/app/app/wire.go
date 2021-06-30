//+build wireinject

package app

import "github.com/google/wire"

func NewApplication() (app *Application, err error) {
	wire.Build(appProviders)
	return
}
