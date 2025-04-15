package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FrancoBarrera99/auth-service/internal/auth"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
	h "github.com/FrancoBarrera99/auth-service/internal/transport/http"
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

	// Routes configuration
	handler := h.NewHandler(serv)
	r := h.NewRouter(serv, *handler)

	addr := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(addr, r.Init()))
}
