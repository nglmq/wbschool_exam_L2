package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONConfig struct {
	Port string `json:"port"`
}

// ReadJSONConfig read config from file
func ReadJSONConfig(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}

	var config JSONConfig

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return "", fmt.Errorf("error decoding json: %v", err)
	}

	return config.Port, nil
}
