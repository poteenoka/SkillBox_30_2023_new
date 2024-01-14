package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP  `yaml:"http"`
	MSSQL `yaml:"mssql"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

type MSSQL struct {
	Instance string `yaml:"instance"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
