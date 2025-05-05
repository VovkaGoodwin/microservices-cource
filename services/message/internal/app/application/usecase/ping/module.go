package ping

import "go.uber.org/fx"

var Module = fx.Module("ping",
	fx.Provide(
		New,
	),
)
