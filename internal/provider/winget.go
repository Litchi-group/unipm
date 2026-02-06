package provider

import (
	"fmt"
	"strings"
)

// WinGetProvider handles WinGet package management
type WinGetProvider struct {
	BaseProvider
}

// NewWinGetProvider creates a new WinGet provider
func NewWinGetProvider() *WinGetProvider {
	return &WinGetProvider{
		BaseProvider: BaseProvider{
			name:       "winget",
			executable: "winget",
		},
	}
}

// Install installs a package using WinGet
func (p *WinGetProvider) Install(spec ProviderSpec) error {
	packageID := spec.ID
	if packageID == "" {
		packageID = spec.Name
	}
	
	args := []string{"install", "--id", packageID, "--silent", "--accept-package-agreements", "--accept-source-agreements"}
	
	fmt.Printf("  → %s\n", FormatCommand("winget", args...))
	return execCommandSilent("winget", args...)
}

// IsInstalled checks if a package is installed
func (p *WinGetProvider) IsInstalled(spec ProviderSpec) bool {
	packageID := spec.ID
	if packageID == "" {
		packageID = spec.Name
	}
	
	args := []string{"list", "--id", packageID}
	
	output, err := execCommand("winget", args...)
	
	// Check if the package ID appears in the output
	return err == nil && strings.Contains(strings.ToLower(output), strings.ToLower(packageID))
}

// InstallCommand returns the command that would be executed
func (p *WinGetProvider) InstallCommand(spec ProviderSpec) string {
	packageID := spec.ID
	if packageID == "" {
		packageID = spec.Name
	}
	
	args := []string{"install", "--id", packageID}
	
	return FormatCommand("winget", args...)
}
