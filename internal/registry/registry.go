package registry

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Litchi-group/unipm/internal/errors"
	"gopkg.in/yaml.v3"
)

const (
	// DefaultRegistryURL is the default GitHub raw URL for the registry
	DefaultRegistryURL = "https://raw.githubusercontent.com/Litchi-group/unipm-registry/main/packages"
	
	// IndexURL is the URL for the package index
	IndexURL = "https://raw.githubusercontent.com/Litchi-group/unipm-registry/main/index.yaml"
	
	// CacheDir is the local cache directory
	CacheDir = ".unipm/cache"
)

// Package represents a package definition from the registry
type Package struct {
	ID           string                       `yaml:"id"`
	Name         string                       `yaml:"name"`
	Homepage     string                       `yaml:"homepage"`
	Dependencies []string                     `yaml:"dependencies,omitempty"`
	Providers    map[string][]ProviderMapping `yaml:"providers"`
}

// ProviderMapping represents OS-specific provider configuration
type ProviderMapping struct {
	Type    string `yaml:"type"`    // "brew", "brew_cask", "winget", "apt", "snap"
	Name    string `yaml:"name"`    // Package name
	ID      string `yaml:"id"`      // Package ID (for winget)
	Classic bool   `yaml:"classic"` // Classic mode (for snap)
}

// PackageInfo represents minimal package information for listing/searching
type PackageInfo struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

// PackageIndex represents the package index file
type PackageIndex struct {
	Packages []PackageInfo `yaml:"packages"`
}

// Registry manages package definitions
type Registry struct {
	baseURL   string
	cacheDir  string
	cacheTTL  time.Duration
	client    *http.Client
}

// NewRegistry creates a new Registry instance
func NewRegistry() *Registry {
	homeDir, _ := os.UserHomeDir()
	cacheDir := filepath.Join(homeDir, CacheDir)
	
	return &Registry{
		baseURL:  DefaultRegistryURL,
		cacheDir: cacheDir,
		cacheTTL: 24 * time.Hour, // Cache for 24 hours
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// LoadPackage loads a package definition by ID
func (r *Registry) LoadPackage(packageID string) (*Package, error) {
	// Try cache first
	if pkg, err := r.loadFromCache(packageID); err == nil {
		return pkg, nil
	}
	
	// Fetch from remote
	pkg, err := r.fetchPackage(packageID)
	if err != nil {
		return nil, err
	}
	
	// Save to cache
	_ = r.saveToCache(packageID, pkg)
	
	return pkg, nil
}

// fetchPackage fetches a package definition from the remote registry
func (r *Registry) fetchPackage(packageID string) (*Package, error) {
	url := fmt.Sprintf("%s/%s.yaml", r.baseURL, packageID)
	
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, errors.NewNetworkError(url, "failed to fetch package", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == 404 {
		return nil, errors.NewNotFoundError(packageID)
	}
	
	if resp.StatusCode != 200 {
		return nil, errors.NewNetworkError(url, fmt.Sprintf("unexpected status code %d", resp.StatusCode), nil)
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewNetworkError(url, "failed to read response", err)
	}
	
	var pkg Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse package definition: %w", err)
	}
	
	return &pkg, nil
}

// loadFromCache loads a package from the local cache
func (r *Registry) loadFromCache(packageID string) (*Package, error) {
	cachePath := r.getCachePath(packageID)
	
	// Check if cache file exists
	info, err := os.Stat(cachePath)
	if err != nil {
		return nil, err
	}
	
	// Check if cache is expired
	if time.Since(info.ModTime()) > r.cacheTTL {
		return nil, fmt.Errorf("cache expired")
	}
	
	// Read cache file
	data, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, err
	}
	
	var pkg Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}
	
	return &pkg, nil
}

// saveToCache saves a package to the local cache
func (r *Registry) saveToCache(packageID string, pkg *Package) error {
	// Ensure cache directory exists
	if err := os.MkdirAll(r.cacheDir, 0755); err != nil {
		return err
	}
	
	cachePath := r.getCachePath(packageID)
	
	data, err := yaml.Marshal(pkg)
	if err != nil {
		return err
	}
	
	return os.WriteFile(cachePath, data, 0644)
}

// getCachePath returns the cache file path for a package
func (r *Registry) getCachePath(packageID string) string {
	return filepath.Join(r.cacheDir, packageID+".yaml")
}

// LoadIndex loads the package index from the registry
func (r *Registry) LoadIndex() ([]PackageInfo, error) {
	resp, err := r.client.Get(IndexURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch index: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code %d for index", resp.StatusCode)
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read index: %w", err)
	}
	
	var index PackageIndex
	if err := yaml.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("failed to parse index: %w", err)
	}
	
	return index.Packages, nil
}
