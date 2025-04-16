package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	if pw == "" {
		return "", fmt.Errorf("trying to hash an empty password")
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

func CheckPassword(pw string, hashedPw string) error {
	if pw == "" || hashedPw == "" {
		return fmt.Errorf("invalid password check")
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
}
