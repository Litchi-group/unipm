package provider

import (
	"fmt"
	"strings"
)

// BrewProvider handles Homebrew package management
type BrewProvider struct {
	BaseProvider
}

// NewBrewProvider creates a new Homebrew provider
func NewBrewProvider() *BrewProvider {
	return &BrewProvider{
		BaseProvider: BaseProvider{
			name:       "brew",
			executable: "brew",
		},
	}
}

// Install installs a package using Homebrew
func (p *BrewProvider) Install(spec ProviderSpec) error {
	args := []string{"install"}
	
	if spec.Type == "brew_cask" {
		args = append(args, "--cask")
	}
	
	args = append(args, spec.Name)
	
	fmt.Printf("  → %s\n", FormatCommand("brew", args...))
	return execCommandSilent("brew", args...)
}

// IsInstalled checks if a package is installed
func (p *BrewProvider) IsInstalled(spec ProviderSpec) bool {
	var args []string
	
	if spec.Type == "brew_cask" {
		args = []string{"list", "--cask", spec.Name}
	} else {
		args = []string{"list", spec.Name}
	}
	
	output, err := execCommand("brew", args...)
	
	// If command succeeds and output is not empty, it's installed
	return err == nil && strings.TrimSpace(output) != ""
}

// InstallCommand returns the command that would be executed
func (p *BrewProvider) InstallCommand(spec ProviderSpec) string {
	args := []string{"install"}
	
	if spec.Type == "brew_cask" {
		args = append(args, "--cask")
	}
	
	args = append(args, spec.Name)
	
	return FormatCommand("brew", args...)
}
