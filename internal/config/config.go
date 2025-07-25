package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           *string `json:"db_url"`
	CurrentUserName *string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting UserHomeDir: %v", err)
	}
	cfgFilePath := filepath.Join(homeDir, configFileName)
	return cfgFilePath, nil
}
