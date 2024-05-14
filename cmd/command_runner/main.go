package main

import (
	"context"
	"log"

	serviceProvider "github.com/Artenso/command-runner/internal/app/service_provider"
)

func main() {
	ctx := context.Background()

	app, err := serviceProvider.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	if err = app.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
