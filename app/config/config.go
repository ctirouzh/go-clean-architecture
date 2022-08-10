package config

import (
	"encoding/json"
	"os"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Server `json:"server"`
}

func LoadConfigs() (*Config, error) {
	fileByte, readFileErr := os.ReadFile("./app/config/config.json")
	if readFileErr != nil {
		return nil, readFileErr
	}

	var config Config
	if err := json.Unmarshal(fileByte, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
