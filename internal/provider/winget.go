package provider

import "strings"

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

// getPackageID returns the package ID or name
func (p *WinGetProvider) getPackageID(spec ProviderSpec) string {
	if spec.ID != "" {
		return spec.ID
	}
	return spec.Name
}

// Install installs a package using WinGet
func (p *WinGetProvider) Install(spec ProviderSpec) error {
	args := p.buildInstallArgs(spec)
	return p.executeWithDisplay(args...)
}

// IsInstalled checks if a package is installed
func (p *WinGetProvider) IsInstalled(spec ProviderSpec) bool {
	packageID := p.getPackageID(spec)
	args := []string{"list", "--id", packageID}

	output, err := execCommand("winget", args...)
	return err == nil && strings.Contains(strings.ToLower(output), strings.ToLower(packageID))
}

// InstallCommand returns the command that would be executed
func (p *WinGetProvider) InstallCommand(spec ProviderSpec) string {
	packageID := p.getPackageID(spec)
	return FormatCommand("winget", "install", "--id", packageID)
}

// Remove removes a package using WinGet
func (p *WinGetProvider) Remove(spec ProviderSpec) error {
	args := p.buildRemoveArgs(spec)
	return p.executeWithDisplay(args...)
}

// RemoveCommand returns the uninstall command
func (p *WinGetProvider) RemoveCommand(spec ProviderSpec) string {
	packageID := p.getPackageID(spec)
	return FormatCommand("winget", "uninstall", "--id", packageID)
}

// buildInstallArgs builds installation arguments
func (p *WinGetProvider) buildInstallArgs(spec ProviderSpec) []string {
	packageID := p.getPackageID(spec)
	return []string{"install", "--id", packageID, "--silent", "--accept-package-agreements", "--accept-source-agreements"}
}

// buildRemoveArgs builds removal arguments
func (p *WinGetProvider) buildRemoveArgs(spec ProviderSpec) []string {
	packageID := p.getPackageID(spec)
	return []string{"uninstall", "--id", packageID, "--silent"}
}
