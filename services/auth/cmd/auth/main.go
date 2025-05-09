package main

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/auth/internal/app"
)

func main() {
	ctx := context.Background()
	container := app.NewMainApp()

	err := container.Start(ctx)
	if err != nil {
		panic(err)
	}
}
