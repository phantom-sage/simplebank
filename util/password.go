package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPasswrod compute and return hashed password.
func HashPasswrod(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	return string(hashedPassword), nil
}

// CheckPassword check if password and hashed password are the same.
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
