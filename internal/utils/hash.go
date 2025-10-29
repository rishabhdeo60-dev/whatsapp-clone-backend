package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	// Generate a hashed password using bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a plaintext password with its hashed version.
func CheckPasswordHash(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
