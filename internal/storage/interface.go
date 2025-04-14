package storage

import (
	m "github.com/FrancoBarrera99/auth-service/internal/auth/model"
)

type UserStorage interface {
	CreateUser(username string, password string, email string) (*m.User, error)
	GetUserByEmail(email string) (*m.User, error)
	Close() error
}
