package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var Conf Config

const DEFAULT_CONFIG_FILENAME = "config.json"

type Config struct {
	GHToken string `json:"GITHUB_TOKEN"`
	AESKey  []byte `json:"AES_KEY"`
}

func (cfg *Config) Load(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	} else {
		f, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(f, cfg)
		if err != nil {
			return err
		}
		return nil
	}
}

func (cfg *Config) Validate() error {
	if cfg.GHToken == "" {
		return errors.New("GITHUB_TOKEN config property is required")
	}
	if len(cfg.AESKey) == 0 {
		return errors.New("AES_KEY config property is required")
	}
	return nil
}

func CreateConfigDirIfRequired() (string, error) {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "g3")
	if err := os.MkdirAll(configDir, 0o700); err != nil {
		return "", err
	}
	return configDir, nil
}
