package config

import (
	"encoding/json"
	"os"
)

func Read() (cfg Config, err error) {
	gatorConfigPath := getConfigFilePath()
	data, err := os.ReadFile(gatorConfigPath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &cfg)
	return
}
