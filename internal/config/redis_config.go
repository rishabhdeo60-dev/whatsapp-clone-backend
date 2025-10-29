package config

import (
	"fmt"
	"os"
)

type RedisConfig struct {
	Addr     string
	Password string
}

func (redisConfig *RedisConfig) Load() error {
	redisConfig.Addr = os.Getenv("REDIS_ADDR")
	redisConfig.Password = os.Getenv("REDIS_PASSWORD")

	if redisConfig.Addr == "" {
		return fmt.Errorf("redis configuration is incomplete")
	}
	return nil
}
