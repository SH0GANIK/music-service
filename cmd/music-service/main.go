package main

//go:generate swag init --parseInternal --dir ./../../internal/handler

import (
	"context"
	"log"
	"music-service/app"
	_ "music-service/docs"
)

// @title Music Service API
// @version 1.0
// @description API server that provides a list of CRUD operations for working with the song structure
// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
