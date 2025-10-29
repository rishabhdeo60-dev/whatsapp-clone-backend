package config

import "os"

type WebSocketConfig struct {
	Port string
}

func (webSocketConfig *WebSocketConfig) Load() error {
	webSocketConfig.Port = os.Getenv("WS_PORT")
	if webSocketConfig.Port == "" {
		webSocketConfig.Port = "8080" // default
	}
	return nil
}
