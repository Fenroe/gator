package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func write(cfg Config) error {
	configFilePath := getConfigFilePath()
	json, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("encountered an error while marshalling config struct to JSON: %v", err)
	}
	mode := int(0777)
	os.WriteFile(configFilePath, json, os.FileMode(mode))
	return nil
}
