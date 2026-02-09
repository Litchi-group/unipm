package provider

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/detector"
)

// GetProvidersForOS returns available providers for the given OS
func GetProvidersForOS(osInfo *detector.OSInfo) []Provider {
	var providers []Provider

	switch {
	case osInfo.IsMacOS():
		providers = []Provider{
			NewBrewProvider(),
		}
	case osInfo.IsWindows():
		providers = []Provider{
			NewWinGetProvider(),
		}
	case osInfo.IsLinux():
		providers = []Provider{
			NewAptProvider(),
			NewSnapProvider(),
		}
	}

	return providers
}

// GetProviderByType returns a provider instance for the given type
func GetProviderByType(providerType string) (Provider, error) {
	switch providerType {
	case "brew", "brew_cask":
		return NewBrewProvider(), nil
	case "winget":
		return NewWinGetProvider(), nil
	case "apt":
		return NewAptProvider(), nil
	case "snap":
		return NewSnapProvider(), nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", providerType)
	}
}
