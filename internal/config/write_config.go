package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = &user
	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("Error setting user: %v", err)
	}
	return nil
}

func write(cfg Config) error {
	cfgFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error - getConfigFilePath: %v", err)
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	outFile, err := os.Create(cfgFilePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer func() {
		err := outFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	_, err = outFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}
