package provider

import (
	"fmt"
	"os/exec"
	"strings"
)

// Provider defines the interface for package managers
type Provider interface {
	// Name returns the provider name (e.g., "brew", "winget")
	Name() string

	// IsAvailable checks if the provider is installed and available
	IsAvailable() bool

	// Install installs a package
	Install(spec ProviderSpec) error

	// Remove removes a package
	Remove(spec ProviderSpec) error

	// IsInstalled checks if a package is already installed
	IsInstalled(spec ProviderSpec) bool

	// InstallCommand returns the command that would be executed
	InstallCommand(spec ProviderSpec) string

	// RemoveCommand returns the uninstall command
	RemoveCommand(spec ProviderSpec) string
}

// ProviderSpec contains provider-specific package information
type ProviderSpec struct {
	Type    string // "brew", "brew_cask", "winget", "apt", "snap"
	Name    string // Package name
	ID      string // Package ID (for winget)
	Classic bool   // Classic mode (for snap)
}

// BaseProvider provides common functionality for all providers
type BaseProvider struct {
	name       string
	executable string
}

// Name returns the provider name
func (p *BaseProvider) Name() string {
	return p.name
}

// IsAvailable checks if the executable is in PATH
func (p *BaseProvider) IsAvailable() bool {
	_, err := exec.LookPath(p.executable)
	return err == nil
}

// execCommand executes a command and returns the output
func execCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

// execCommandSilent executes a command and returns only the error
func execCommandSilent(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

// FormatCommand formats a command for display
func FormatCommand(name string, args ...string) string {
	parts := append([]string{name}, args...)
	return strings.Join(parts, " ")
}

// GetInstallationGuide returns installation instructions for missing providers
func GetInstallationGuide(providerName string) string {
	guides := map[string]string{
		"brew": `Homebrew is not installed.
Install it from: https://brew.sh

  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`,
		
		"winget": `WinGet is not installed.
Install it from: https://aka.ms/getwinget

WinGet comes pre-installed on Windows 11 and recent Windows 10 builds.
If missing, install "App Installer" from the Microsoft Store.`,
		
		"apt": `APT is not installed.
APT comes pre-installed on Debian-based systems (Ubuntu, Debian).
If you're not on a Debian-based system, unipm may not support your distribution yet.`,
		
		"snap": `Snap is not installed.
Install it with:

  sudo apt update
  sudo apt install snapd`,
	}

	if guide, ok := guides[providerName]; ok {
		return guide
	}

	return fmt.Sprintf("No installation guide available for %s", providerName)
}
