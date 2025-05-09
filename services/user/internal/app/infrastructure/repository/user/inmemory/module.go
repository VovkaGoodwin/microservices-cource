package inmemory

import "go.uber.org/fx"

var Module = fx.Module(
	"user-inmemory-repository",
	fx.Provide(New),
)
