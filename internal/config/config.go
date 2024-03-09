package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type PGConfig struct {
	PGUser     string `env:"POSTGRES_USER"`
	PGPassword string `env:"POSTGRES_PASSWORD"`
	PGPort     string `env:"POSTGRES_CONTAINER_PORT"`
	PGHost     string `env:"POSTGRES_HOST"`
	PGDatabase string `env:"POSTGRES_DB"`
	PGSSLMode  string `env:"POSTGRES_SSL_MODE"`
}

type MinioConfig struct {
	Host     string `env:"MINIO_HOST"`
	Port     string `env:"MINIO_PORT"`
	User     string `env:"MINIO_USER"`
	Password string `env:"MINIO_PASSWORD"`
	SSL      bool   `env:"MINIO_SSL"`
}

type RedisConfig struct {
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}

type Config struct {
	PGConfig
	RedisConfig
	MinioConfig

	Salt     string `env:"SALT"`
	SignKey  string `env:"SIGN_KEY"`
	HTTPPort int    `env:"HTTP_PORT"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadEnv(config); err != nil {
			log.Fatalf("failed to parse configs: %v", err)
		}
	})

	return config, nil
}
