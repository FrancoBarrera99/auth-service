package token

import (
	"time"

	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID       string
	Username string
	Email    string
	Password string
	jwt.RegisteredClaims
}

func GenerateJWT(user *model.User, secret string) (string, error) {
	claims := Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
