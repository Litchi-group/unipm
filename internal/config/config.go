package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// DevPack represents the devpack.yaml configuration
type DevPack struct {
	Apps     []string              `yaml:"apps"`
	Profiles map[string][]string   `yaml:"profiles,omitempty"`
}

// PackageSpec represents a package with optional version
type PackageSpec struct {
	Name    string
	Version string
}

// ParsePackageSpec parses a package spec (e.g., "node@18.x", "git")
func ParsePackageSpec(spec string) PackageSpec {
	parts := strings.SplitN(spec, "@", 2)
	
	if len(parts) == 2 {
		return PackageSpec{
			Name:    parts[0],
			Version: parts[1],
		}
	}
	
	return PackageSpec{
		Name: parts[0],
	}
}

// GetApps returns the list of apps, optionally filtered by profile
func (d *DevPack) GetApps(profile string) []string {
	if profile == "" {
		return d.Apps
	}
	
	if profileApps, ok := d.Profiles[profile]; ok {
		return profileApps
	}
	
	return d.Apps
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
