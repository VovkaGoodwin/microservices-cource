package main

import (
	"context"

	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app"
)

func main() {
	container := app.NewMainApp()
	err := container.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
