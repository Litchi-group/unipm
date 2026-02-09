package registry

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/provider"
)

// Resolver resolves package IDs to provider specifications
type Resolver struct {
	registry *Registry
	osInfo   *detector.OSInfo
}

// NewResolver creates a new Resolver
func NewResolver(registry *Registry, osInfo *detector.OSInfo) *Resolver {
	return &Resolver{
		registry: registry,
		osInfo:   osInfo,
	}
}

// Resolve resolves a package ID to a provider specification
func (r *Resolver) Resolve(packageID string) (*provider.ProviderSpec, error) {
	// Load package definition
	pkg, err := r.registry.LoadPackage(packageID)
	if err != nil {
		return nil, err
	}

	// Get OS-specific providers
	osKey := r.getOSKey()
	mappings, ok := pkg.Providers[osKey]
	if !ok || len(mappings) == 0 {
		return nil, fmt.Errorf("no provider available for %s on %s", packageID, osKey)
	}

	// Use the first available provider
	// TODO: Implement provider priority/fallback
	mapping := mappings[0]

	return &provider.ProviderSpec{
		Type:    mapping.Type,
		Name:    mapping.Name,
		ID:      mapping.ID,
		Classic: mapping.Classic,
	}, nil
}

// getOSKey returns the OS key for provider lookup
func (r *Resolver) getOSKey() string {
	switch {
	case r.osInfo.IsMacOS():
		return "macos"
	case r.osInfo.IsWindows():
		return "windows"
	case r.osInfo.IsLinux():
		return "linux"
	default:
		return "unknown"
	}
}

// ResolveAll resolves multiple package IDs
func (r *Resolver) ResolveAll(packageIDs []string) (map[string]*provider.ProviderSpec, error) {
	result := make(map[string]*provider.ProviderSpec)

	for _, id := range packageIDs {
		spec, err := r.Resolve(id)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve %s: %w", id, err)
		}
		result[id] = spec
	}

	return result, nil
}
