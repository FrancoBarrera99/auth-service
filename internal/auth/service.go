package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/FrancoBarrera99/auth-service/internal/auth/strategies/local"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
	"github.com/FrancoBarrera99/auth-service/internal/utils"
	"github.com/FrancoBarrera99/auth-service/internal/utils/token"
)

type Service struct {
	stor      storage.UserStorage
	strats    map[string]AuthStrategy
	jwtSecret string
}

func NewService(stor storage.UserStorage) (*Service, error) {
	secret := os.Getenv("JWT_SECRET")
	strats := map[string]AuthStrategy{
		"local": local.NewLocalAuth(stor),
		//"google": oauth.NewGoogleAuth(stor),
	}

	return &Service{stor: stor, strats: strats, jwtSecret: secret}, nil
}

func (s *Service) Login(creds model.Credentials) (*model.User, string, error) {
	strat, ok := s.strats[creds.Method]
	if !ok {
		return nil, "", fmt.Errorf("unsupported method %s", creds.Method)
	}

	user, err := strat.ValidateCredentials(creds.Data)
	if err != nil {
		return nil, "", err
	}

	tkn, err := token.GenerateJWT(user.ID, s.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return user, tkn, nil
}

func (s *Service) Register(reg model.UserRegister) (string, error) {
	if reg.Username == "" || reg.Password == "" || reg.Email == "" {
		return "", fmt.Errorf("all fields are required for register a new user")
	}
	if !strings.Contains(reg.Email, "@") {
		return "", fmt.Errorf("invalid email format")
	}
	if len(reg.Password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters long")
	}

	hashedPw, err := utils.HashPassword(reg.Password)
	if err != nil {
		return "", err
	}

	user, err := s.stor.CreateUser(reg.Username, hashedPw, reg.Email)
	if err != nil {
		return "", err
	}

	return token.GenerateJWT(user.ID, s.jwtSecret)
}

func (s *Service) ValidateToken(tokenString string) (bool, error) {
	_, err := token.ValidateJWT(tokenString, s.jwtSecret)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) GetAuthURL(method string, state string) (string, error) {
	return "", nil
}
