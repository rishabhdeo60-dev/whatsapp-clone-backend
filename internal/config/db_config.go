package config

// Todo: Add database configuration related code here in the future.

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (dbConfig *DBConfig) Load() error {

	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port = os.Getenv("DB_PORT")
	dbConfig.User = os.Getenv("DB_USER")
	dbConfig.Password = os.Getenv("DB_PASS")
	dbConfig.Name = os.Getenv("DB_NAME")
	dbConfig.SSLMode = os.Getenv("DB_SSLMODE")

	if dbConfig.Host == "" || dbConfig.Port == "" || dbConfig.User == "" || dbConfig.Password == "" || dbConfig.Name == "" || dbConfig.SSLMode == "" {
		return fmt.Errorf("database configuration is incomplete")
	}
	return nil
}
