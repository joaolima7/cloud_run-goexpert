package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WebServerPort string
	WeatherKey    string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		WebServerPort: getEnv("WEB_SERVER_PORT"),
		WeatherKey:    getEnv("WEATHER_API_KEY"),
	}, nil
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	panic("erro ao ler a variavel de ambiente: " + key)
}
