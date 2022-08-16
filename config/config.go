package config

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrFailedToReadFile  = errors.New("failed to read config file")
	ErrFailedToUnmarshal = errors.New("failed to unmarshal config file")
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Server `json:"server"`
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

	return &config, nil
}
