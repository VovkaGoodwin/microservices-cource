package http_server

import "go.uber.org/fx"

var Module = fx.Module(
	"http",
	fx.Provide(
		NewHandler,
		NewHttpServer,
	),
)
