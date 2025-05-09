package main

import (
	"context"

	_ "github.com/VovkaGoodwin/microservices-cource/services/gateway/cmd/docs"
	"github.com/VovkaGoodwin/microservices-cource/services/gateway/internal/app"
)

// @BasePath /api
// @host localhost:80
func main() {
	ctx := context.Background()
	mainApp := app.NewMainApp()

	err := mainApp.Start(ctx)
	if err != nil {
		panic(err)
	}
}
