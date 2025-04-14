package oauth

import (
	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
)

type GoogleAuth struct {
	stor storage.UserStorage
}

func NewGoogleAuth(stor storage.UserStorage) *GoogleAuth {
	return &GoogleAuth{stor: stor}
}

func (g *GoogleAuth) Validate(creds model.Credentials) (*model.User, string, error) {
	return nil, "", nil
}

func (l *GoogleAuth) GetAuthURL(method string, state string) (string, error) {
	return "", nil
}
