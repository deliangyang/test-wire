package app

import "github.com/google/wire"

var appProviders = wire.NewSet(
	wire.Struct(new(Application), "*"),
	testProvider,
)

var testProvider = wire.NewSet(
	wire.Struct(new(Test), "*"),
)
