package config

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type AppConfig struct {
	DBConfig    *DBConfig
	RedisConfig *RedisConfig
	KafkaConfig *KafkaConfig
	WSConfig    *WebSocketConfig
}

func (config *AppConfig) LoadConfig() *AppConfig {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	if err := config.DBConfig.Load(); err != nil {
		log.Fatal("Failed to load database configuration:", err)
	}

	if err := config.RedisConfig.Load(); err != nil {
		log.Fatal("Failed to load redis configuration:", err)
	}

	if err := config.KafkaConfig.Load(); err != nil {
		log.Fatal("Failed to load kafka configuration:", err)
	}

	if err := config.WSConfig.Load(); err != nil {
		log.Fatal("Failed to load websocket configuration:", err)
	}

	return &AppConfig{
		DBConfig:    config.DBConfig,
		RedisConfig: config.RedisConfig,
		KafkaConfig: config.KafkaConfig,
		WSConfig:    config.WSConfig,
	}
}
