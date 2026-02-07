package provider

import (
	"fmt"
	"strings"
)

// SnapProvider handles Snap package management
type SnapProvider struct {
	BaseProvider
}

// NewSnapProvider creates a new Snap provider
func NewSnapProvider() *SnapProvider {
	return &SnapProvider{
		BaseProvider: BaseProvider{
			name:       "snap",
			executable: "snap",
		},
	}
}

// Install installs a package using Snap
func (p *SnapProvider) Install(spec ProviderSpec) error {
	args := []string{"install", spec.Name}
	
	if spec.Classic {
		args = append(args, "--classic")
	}
	
	fmt.Printf("  → %s\n", FormatCommand("sudo snap", args...))
	
	// Snap requires sudo
	return execCommandSilent("sudo", append([]string{"snap"}, args...)...)
}

// IsInstalled checks if a package is installed
func (p *SnapProvider) IsInstalled(spec ProviderSpec) bool {
	args := []string{"list", spec.Name}
	
	output, err := execCommand("snap", args...)
	
	// Check if package name appears in the list
	if err != nil {
		return false
	}
	
	for _, line := range strings.Split(output, "\n") {
		fields := strings.Fields(line)
		if len(fields) > 0 && fields[0] == spec.Name {
			return true
		}
	}
	
	return false
}

// InstallCommand returns the command that would be executed
func (p *SnapProvider) InstallCommand(spec ProviderSpec) string {
	args := []string{"install", spec.Name}
	
	if spec.Classic {
		args = append(args, "--classic")
	}
	
	return FormatCommand("sudo snap", args...)
}

// Remove removes a package using Snap
func (p *SnapProvider) Remove(spec ProviderSpec) error {
	args := []string{"remove", spec.Name}
	
	fmt.Printf("  → %s\n", FormatCommand("sudo snap", args...))
	
	// Snap requires sudo
	return execCommandSilent("sudo", append([]string{"snap"}, args...)...)
}

// RemoveCommand returns the uninstall command
func (p *SnapProvider) RemoveCommand(spec ProviderSpec) string {
	args := []string{"remove", spec.Name}
	
	return FormatCommand("sudo snap", args...)
}
