package main

import (
	"log"

	"github.com/FrancoBarrera99/auth-service/internal/auth"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
)

func main() {
	// Storage configuration
	stor, err := storage.CreateStorage()
	if err != nil {
		log.Fatal(err)
	}

	defer stor.Close()

	// Service configuration
	serv, err := auth.NewService(stor)
	if err != nil {
		log.Fatal(err)
	}
}
