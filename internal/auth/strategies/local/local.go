package local

import (
	"fmt"

	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
	"github.com/FrancoBarrera99/auth-service/internal/utils"
)

type LocalAuth struct {
	stor storage.UserStorage
}

func NewLocalAuth(stor storage.UserStorage) *LocalAuth {
	return &LocalAuth{stor: stor}
}

func (l *LocalAuth) ValidateCredentials(creds map[string]interface{}) (*model.User, error) {
	email, ok := creds["email"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid email provided")
	}

	pw, ok := creds["password"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid password provided")
	}

	// Get user from storage
	user, uErr := l.stor.GetUserByEmail(email)
	if uErr != nil {
		return nil, uErr
	}

	pwErr := utils.CheckPassword(pw, user.Password)
	if pwErr != nil {
		return nil, pwErr
	}

	return user, nil
}

func (l *LocalAuth) GetAuthURL(method string, state string) (string, error) {
	return "", nil
}
