package config

import (
	"fmt"
	"vk-test-task/internal/models"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server    ServerConfig
	Tarantool TarantoolConfig
}

type ServerConfig struct {
	Mode   string `env:"APP_MODE" env-default:"local"`
	Domain string `env:"DOMAIN" env-default:"localhost"`
}

type TarantoolConfig struct {
	Host     string `env:"TARANTOOL_HOST"`
	Port     string `env:"TARANTOOL_PORT"`
	User     string `env:"TARANTOOL_USER"`
	Password string `env:"TARANTOOL_PASSWORD"`
}

func New() (*Config, error) {
	var (
		err error
		cfg Config
	)

	err = cleanenv.ReadConfig(".env", &cfg)

	if cfg == (Config{}) {
		return nil, models.ErrEmptyConfig
	}

	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
