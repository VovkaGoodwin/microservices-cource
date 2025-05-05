package main

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app"
)

func main() {
	container := app.NewMainApp()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "app", "user")
	err := container.Start(ctx)
	if err != nil {
		panic(err)
	}
}
