package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// DevPack represents the devpack.yaml configuration
type DevPack struct {
	Apps []string `yaml:"apps"`
}

// Load loads the devpack.yaml configuration from the given path
func Load(path string) (*DevPack, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", path, err)
	}
	
	var config DevPack
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", path, err)
	}
	
	return &config, nil
}
