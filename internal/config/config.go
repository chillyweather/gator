package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath(filename string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}
	filePath := homeDir + "/" + filename
	return filePath, nil
}

func Read() (Config, error) {
	var configStruct Config

	filePath, err := getConfigFilePath(configFileName)
	if err != nil {
		return configStruct, fmt.Errorf("Error: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return configStruct, fmt.Errorf("Error: %v", err)
	}

	if err := json.Unmarshal(data, &configStruct); err != nil {
		return configStruct, fmt.Errorf("Error: %v", err)
	}

	return configStruct, nil
}

func (c Config) SetUser(userName string) error {

	filePath, err := getConfigFilePath(configFileName)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	c.CurrentUserName = userName

	modData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	err = os.WriteFile(filePath, modData, 0644)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	return nil
}
