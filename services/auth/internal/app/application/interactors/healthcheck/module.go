package healthcheck

import "go.uber.org/fx"

var Module = fx.Module(
	"healthcheck",
	fx.Provide(NewInteractor),
)
