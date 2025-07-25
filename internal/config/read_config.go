package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	var cfg Config

	cfgFilePath, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("Error - getConfigFilePath: %v", err)
	}

	content, err := os.ReadFile(cfgFilePath)
	if err != nil {
		return cfg, fmt.Errorf("Error reading file: %v", err)
	}

	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	return cfg, nil
}
