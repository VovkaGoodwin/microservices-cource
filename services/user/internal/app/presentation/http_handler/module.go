package http_handler

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/healthcheck"
	"github.com/VovkaGoodwin/microservices-cource/pkg/http_server"
	"go.uber.org/fx"
)

var Module = fx.Module("http-server", fx.Provide(
	func(hc healthcheck.Interactor) *Deps {
		return &Deps{
			hc: hc,
		}
	},
	NewHttpHandler,
),
	http_server.Module,
)
