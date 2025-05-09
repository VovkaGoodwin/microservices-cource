package http_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/application/interactors/healthcheck"
	"go.uber.org/fx"
)

var Module = fx.Module("http-server", fx.Provide(
	func(hc healthcheck.Interactor) *Deps {
		return &Deps{
			hc: hc,
		}
	},
	NewHttpHandler,
	NewHttpServer,
))
