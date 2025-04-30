package main

import (
	"auth/internal/app"
	"auth/internal/config"
	"context"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	mainApp := app.New(ctx, cfg)

	mainApp.Start(ctx, cfg)
}
