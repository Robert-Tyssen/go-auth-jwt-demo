package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Receives a password as an input and returns a hash that can be stored in the database
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compares a password with a hash to determine if the password is correct
func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Validates that a password meets complexity requirements
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}
	return nil
}
