package main

import (
	"log"

	"github.com/FrancoBarrera99/auth-service/internal/storage"
)

func main() {
	// Repository configuration
	repo, err := storage.CreateStorage()
	if err != nil {
		log.Fatal(err)
	}

	defer repo.Close()
}
