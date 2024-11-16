package main

import (
	"auth/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app")
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app")
	}
}
