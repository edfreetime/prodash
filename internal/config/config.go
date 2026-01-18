package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Project struct {
	Name     string            `josn:"name"`
	Path     string            `json:"path"`
	Type     string            `json:"type"`
	Commands map[string]string `json:"commands"`
}

type Config struct {
	Projects []Project `json:"projects"`
}

func getConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "prodash")

	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return "", err
	}

	return filepath.Join(appDir, "config.json"), nil
}

func defaultConfig() Config {
	return Config{
		Projects: []Project{
			{
				Name: "Example Project",
				Path: "~/project",
				Type: "go",
				Commands: map[string]string{
					"open": "nv .",
				},
			},
		},
	}
}

func Load() (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	// create default if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cfg := defaultConfig()
		if err := Save(cfg); err != nil {
			return Config{}, err
		}
		return cfg, nil
	}

	// read file
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func Save(cfg Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0o644)
}
