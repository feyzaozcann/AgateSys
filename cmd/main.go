package main

import (
	"context"
	"log"

	"agate-project/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv, err := server.NewServer(ctx)
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}
	defer srv.Close()

	if err := srv.Run("localhost:8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
