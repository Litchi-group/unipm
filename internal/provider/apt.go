package provider

import (
	"fmt"
	"strings"
)

// AptProvider handles APT package management
type AptProvider struct {
	BaseProvider
}

// NewAptProvider creates a new APT provider
func NewAptProvider() *AptProvider {
	return &AptProvider{
		BaseProvider: BaseProvider{
			name:       "apt",
			executable: "apt",
		},
	}
}

// Install installs a package using APT
func (p *AptProvider) Install(spec ProviderSpec) error {
	args := []string{"install", "-y", spec.Name}
	
	fmt.Printf("  → %s\n", FormatCommand("sudo apt", args...))
	
	// APT requires sudo
	return execCommandSilent("sudo", append([]string{"apt"}, args...)...)
}

// IsInstalled checks if a package is installed
func (p *AptProvider) IsInstalled(spec ProviderSpec) bool {
	args := []string{"list", "--installed", spec.Name}
	
	output, err := execCommand("apt", args...)
	
	// Check if package name appears in installed list
	if err != nil {
		return false
	}
	
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, spec.Name+"/") {
			return true
		}
	}
	
	return false
}

// InstallCommand returns the command that would be executed
func (p *AptProvider) InstallCommand(spec ProviderSpec) string {
	args := []string{"install", "-y", spec.Name}
	return FormatCommand("sudo apt", args...)
}
