package provider

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
	args := p.buildInstallArgs(spec)
	return p.executeWithDisplay(args...)
}

// IsInstalled checks if a package is installed
func (p *BrewProvider) IsInstalled(spec ProviderSpec) bool {
	args := []string{"list"}
	
	if spec.Type == "brew_cask" {
		args = append(args, "--cask")
	}
	
	args = append(args, spec.Name)
	return p.checkInstalled(args...)
}

// buildInstallArgs builds installation arguments
func (p *BrewProvider) buildInstallArgs(spec ProviderSpec) []string {
	args := []string{"install"}
	
	if spec.Type == "brew_cask" {
		args = append(args, "--cask")
	}
	
	return append(args, spec.Name)
}

// InstallCommand returns the command that would be executed
func (p *BrewProvider) InstallCommand(spec ProviderSpec) string {
	args := p.buildInstallArgs(spec)
	return FormatCommand("brew", args...)
}

// Remove removes a package using Homebrew
func (p *BrewProvider) Remove(spec ProviderSpec) error {
	args := p.buildRemoveArgs(spec)
	return p.executeWithDisplay(args...)
}

// RemoveCommand returns the uninstall command
func (p *BrewProvider) RemoveCommand(spec ProviderSpec) string {
	args := p.buildRemoveArgs(spec)
	return FormatCommand("brew", args...)
}

// buildRemoveArgs builds removal arguments
func (p *BrewProvider) buildRemoveArgs(spec ProviderSpec) []string {
	args := []string{"uninstall"}
	
	if spec.Type == "brew_cask" {
		args = append(args, "--cask")
	}
	
	return append(args, spec.Name)
}
