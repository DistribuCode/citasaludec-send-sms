package config

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret string
	RedisHost string
	Port      string
}

func LoadConfig() *Config {
	c := &Config{
		JWTSecret: getEnv("JWT_SECRET", "defaultsecret"),
		RedisHost: getEnv("REDIS_HOST", "localhost:6379"),
		Port:      getEnv("PORT", "5000"),
	}
	return c
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("⚠️  Env %s not set, using default %s", key, fallback)
	return fallback
}
