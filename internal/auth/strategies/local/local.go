package local

import (
	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
)

type LocalAuth struct {
	stor storage.UserStorage
}

func NewLocalAuth(stor storage.UserStorage) *LocalAuth {
	return &LocalAuth{stor: stor}
}

func (l *LocalAuth) ValidateCredentials(creds model.Credentials) (*model.User, string, error) {
	return nil, "", nil
}

func (l *LocalAuth) GetAuthURL(method string, state string) (string, error) {
	return "", nil
}
