package redis

import "go.uber.org/fx"

var Module = fx.Module(
	"token-redis-repository",
	fx.Provide(
		newRepository,
	),
)
