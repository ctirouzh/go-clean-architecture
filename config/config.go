package config

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

var (
	ErrFailedToReadFile  = errors.New("failed to read config file")
	ErrFailedToUnmarshal = errors.New("failed to unmarshal config file")
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type JWT struct {
	SecretKey string        `json:"secret_key"`
	TTL       time.Duration `json:"ttl_minute"`
}

type Postgres struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Config struct {
	Server   `json:"server"`
	JWT      `json:"jwt"`
	Postgres `json:"postgres"`
}

func Parse(path string) (*Config, error) {
	file, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, ErrFailedToReadFile
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, ErrFailedToUnmarshal
	}
	config.JWT.TTL *= time.Minute
	return &config, nil
}
