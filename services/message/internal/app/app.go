package app

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"go.uber.org/fx"
)

func NewMainApp() *fx.App {
	return app.InitApp(
		initApplicationLayer(),
		initPresentationLayer(),
		initRunners(),
	)
}
