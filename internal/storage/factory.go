package storage

import (
	"fmt"
	"os"

	"github.com/FrancoBarrera99/auth-service/internal/storage/sql/postgres"
)

func CreateStorage() (UserStorage, error) {
	storage := os.Getenv("STORAGE_TYPE")
	switch storage {
	case "postgres":
		return postgres.NewPostgresStorage()
	default:
		return nil, fmt.Errorf("no valid storage type was found")
	}
}
