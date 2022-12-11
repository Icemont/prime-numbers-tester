package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

var (
	Server ServerConfig
)

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}

	configServer()
}

func configServer() {
	requestTimeout, _ := strconv.Atoi(getEnv("APP_READ_TIMEOUT", "30"))
	httpPort, _ := strconv.Atoi(getEnv("APP_HTTP_PORT", "8888"))

	Server = ServerConfig{
		HttpPort:    httpPort,
		ReadTimeout: time.Duration(requestTimeout) * time.Second,
		AppDomain:   getEnv("APP_DOMAIN", "localhost"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
