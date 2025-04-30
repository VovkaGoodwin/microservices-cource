package main

import (
	"context"

	_ "gateway/cmd/docs"
	"gateway/internal/app"
)

// @BasePath /api
// @host localhost:80

func main() {
	ctx := context.Background()
	mainApp := app.New()

	mainApp.Start(ctx)
}
