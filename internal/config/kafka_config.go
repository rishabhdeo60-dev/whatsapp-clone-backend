package config

import "os"

type KafkaConfig struct {
	Broker string
	Topic  string
}

func (kafkaConfig *KafkaConfig) Load() error {
	kafkaConfig.Broker = os.Getenv("KAFKA_BROKER")
	kafkaConfig.Topic = os.Getenv("KAFKA_TOPIC")
	return nil
}
