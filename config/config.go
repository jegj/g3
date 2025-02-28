package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var Conf Config

type Config struct {
	GHToken  string
	LogLevel string
}

func (cfg *Config) Load(filePath string) error {
	_, err := os.Stat(filePath)
	configFileExists := err == nil
	if configFileExists {
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

		_, err := os.Create(filePath)
		if err != nil {
			return err
		}

		return nil
	}
}

func GetDefaultConfigDir() (dir string, err error) {
	if env, ok := os.LookupEnv("G3_CONFIG_DIR"); ok {
		dir = env
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config", "g3")
	}

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("cannot create config directory: %v", err)
	}

	return dir, nil
}
