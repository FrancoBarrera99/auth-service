package auth

import (
	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
)

type AuthService interface {
	Login(creds model.Credentials) (*model.User, string, error)
	Register(reg model.UserRegister) (string, error)
	ValidateToken(tokenString string) (bool, error)
	GetAuthURL(method string, state string) (string, error)
}

type AuthStrategy interface {
	ValidateCredentials(creds model.Credentials) (*model.User, string, error)
	GetAuthURL(method string, state string) (string, error)
}
