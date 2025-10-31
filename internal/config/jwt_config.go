package config

import (
	"os"
	"strconv"
)

type JWTConfig struct {
	JwtSecret     string
	TokenDuration int // in minutes
}

func (jwtConfig *JWTConfig) Load() error {
	// Load JWT configuration from environment variables
	jwtConfig.JwtSecret = os.Getenv("JWT_SECRET")
	durationStr := os.Getenv("JWT_EXPIRATION_HOURS")
	if duration, err := strconv.Atoi(durationStr); err == nil {
		jwtConfig.TokenDuration = duration
	}
	return nil
}
