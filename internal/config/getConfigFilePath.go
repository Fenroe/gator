package config

import (
	"fmt"
	"os"
)

func getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	gatorConfigPath := fmt.Sprintf("%s/.gatorconfig.json", home)
	return gatorConfigPath
}
