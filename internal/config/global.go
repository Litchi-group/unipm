package config

import (
	"os"
	"path/filepath"

	"github.com/Litchi-group/unipm/internal/logger"
	"gopkg.in/yaml.v3"
)

// GlobalConfig represents the ~/.unipm/config.yaml file
type GlobalConfig struct {
	Registry RegistryConfig `yaml:"registry"`
	Log      LogConfig      `yaml:"log"`
}

// RegistryConfig contains registry settings
type RegistryConfig struct {
	URL      string `yaml:"url"`       // Custom registry URL
	CacheTTL int    `yaml:"cache_ttl"` // Cache TTL in hours (default: 24)
}

// LogConfig contains logging settings
type LogConfig struct {
	Level string `yaml:"level"` // debug, info, warn, error
}

// DefaultGlobalConfig returns the default configuration
func DefaultGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		Registry: RegistryConfig{
			URL:      "https://raw.githubusercontent.com/Litchi-group/unipm-registry/main/packages",
			CacheTTL: 24,
		},
		Log: LogConfig{
			Level: "info",
		},
	}
}

// LoadGlobalConfig loads the global configuration from ~/.unipm/config.yaml
func LoadGlobalConfig() (*GlobalConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Warn("Failed to get home directory: %v", err)
		return DefaultGlobalConfig(), nil
	}
	
	configPath := filepath.Join(homeDir, ".unipm", "config.yaml")
	
	// If config file doesn't exist, return default
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Debug("Config file not found, using defaults: %s", configPath)
		return DefaultGlobalConfig(), nil
	}
	
	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		logger.Warn("Failed to read config file: %v", err)
		return DefaultGlobalConfig(), nil
	}
	
	// Parse YAML
	var config GlobalConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		logger.Warn("Failed to parse config file: %v", err)
		return DefaultGlobalConfig(), nil
	}
	
	// Apply defaults for missing fields
	if config.Registry.URL == "" {
		config.Registry.URL = DefaultGlobalConfig().Registry.URL
	}
	if config.Registry.CacheTTL == 0 {
		config.Registry.CacheTTL = DefaultGlobalConfig().Registry.CacheTTL
	}
	if config.Log.Level == "" {
		config.Log.Level = DefaultGlobalConfig().Log.Level
	}
	
	logger.Debug("Loaded config from %s", configPath)
	return &config, nil
}

// Save saves the configuration to ~/.unipm/config.yaml
func (c *GlobalConfig) Save() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	
	configDir := filepath.Join(homeDir, ".unipm")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}
	
	configPath := filepath.Join(configDir, "config.yaml")
	
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	
	return os.WriteFile(configPath, data, 0644)
}
