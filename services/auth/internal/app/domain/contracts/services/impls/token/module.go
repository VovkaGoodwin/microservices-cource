package token

import "go.uber.org/fx"

var Module = fx.Module("token-service", fx.Provide(New))
