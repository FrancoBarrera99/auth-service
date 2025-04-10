package storage

import (
	m "github.com/FrancoBarrera99/auth-service/internal/auth/model"
)

type UserStorage interface {
	SaveUser(user m.User) error
	GetUserByEmail(email string) (*m.User, error)
	Close() error
}
