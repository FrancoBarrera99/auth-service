package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/FrancoBarrera99/auth-service/internal/auth/strategies/local"
	"github.com/FrancoBarrera99/auth-service/internal/auth/strategies/oauth"
	"github.com/FrancoBarrera99/auth-service/internal/storage"
	"github.com/FrancoBarrera99/auth-service/internal/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	stor      storage.UserStorage
	strats    map[string]AuthStrategy
	jwtSecret string
}

func NewService(stor storage.UserStorage) (*Service, error) {
	secret := os.Getenv("JWT_SECRET")
	strats := map[string]AuthStrategy{
		"local":  local.NewLocalAuth(stor),
		"google": oauth.NewGoogleAuth(stor),
	}

	return &Service{stor: stor, strats: strats, jwtSecret: secret}, nil
}

func (s *Service) Login(creds model.Credentials) (*model.User, string, error) {
	strat, ok := s.strats[creds.Method]
	if !ok {
		return nil, "", fmt.Errorf("unsupported method %s", creds.Method)
	}

	user, _, err := strat.Validate(creds)
	if err != nil {
		return nil, "", err
	}

	tkn, err := token.GenerateJWT(user.ID, s.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return user, tkn, nil
}

func (s *Service) Register(username string, password string, email string) (string, error) {
	if username == "" || password == "" || email == "" {
		return "", fmt.Errorf("all fields are required for register a new user")
	}
	if !strings.Contains(email, "@") {
		return "", fmt.Errorf("invalid email format")
	}
	if len(password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters long")
	}

	hashedPw, err := s.hashPassword(password)
	if err != nil {
		return "", err
	}

	user, err := s.stor.CreateUser(username, hashedPw, email)
	if err != nil {
		return "", err
	}

	return token.GenerateJWT(user.ID, s.jwtSecret)
}

func (s *Service) ValidateToken(token string) (*token.Claims, error) {
	return nil, nil
}

func (s *Service) GetAuthURL(method string, state string) (string, error) {
	return "", nil
}

func (s *Service) hashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("")
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password")
	}
	return string(hashedPw), nil
}
