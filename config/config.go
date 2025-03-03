package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var Conf Config

type Config struct {
	GHToken string `json:"GITHUB_TOKEN"`
}

func (cfg *Config) Load(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		f, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(f, cfg)
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}

func CreateConfigDirIfRequired() (string, error) {
	configDir := getG3ConfigDir()
	err := createDir(configDir)
	if err != nil {
		return "", err
	}

	return configDir, nil
}

func createDir(dir string) error {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	return nil
}

func getG3ConfigDir() (dir string) {
	if env, ok := os.LookupEnv("G3_CONFIG_DIR"); ok {
		dir = env
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config", "g3")
	}

	return dir
}
