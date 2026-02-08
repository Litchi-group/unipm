package cmd

import (
	"fmt"
	"os"

	"github.com/Litchi-group/unipm/internal/config"
	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/provider"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var exportCmd = &cobra.Command{
	Use:   "export [file]",
	Short: "Export currently installed packages to devpack.yaml",
	Long: `Scans the system for installed packages and exports them to a devpack.yaml file.

This creates a portable configuration that can be used to recreate the environment.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		outputFile := "devpack.yaml"
		if len(args) > 0 {
			outputFile = args[0]
		}
		return runExport(outputFile)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

func runExport(outputFile string) error {
	// Detect OS
	osInfo := detector.DetectOS()
	
	fmt.Printf("Scanning installed packages on %s...\n\n", osInfo.String())
	
	// Get available providers
	providers := getAvailableProviders(osInfo)
	
	if len(providers) == 0 {
		return fmt.Errorf("no package managers found on this system")
	}
	
	// Collect installed packages
	allPackages := make(map[string]bool)
	
	for _, p := range providers {
		fmt.Printf("Scanning %s packages...\n", p.Name())
		
		packages, err := p.ListInstalled()
		if err != nil {
			fmt.Printf("  Warning: failed to list %s packages: %v\n", p.Name(), err)
			continue
		}
		
		fmt.Printf("  Found %d packages\n", len(packages))
		
		for _, pkg := range packages {
			allPackages[pkg] = true
		}
	}
	
	if len(allPackages) == 0 {
		fmt.Println("\nNo packages found.")
		return nil
	}
	
	// Convert to sorted list
	var packageList []string
	for pkg := range allPackages {
		packageList = append(packageList, pkg)
	}
	
	// Create devpack
	devpack := config.DevPack{
		Apps: packageList,
	}
	
	// Marshal to YAML
	data, err := yaml.Marshal(&devpack)
	if err != nil {
		return fmt.Errorf("failed to generate YAML: %w", err)
	}
	
	// Write to file
	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	fmt.Printf("\n✅ Exported %d packages to %s\n", len(packageList), outputFile)
	
	return nil
}

func getAvailableProviders(osInfo *detector.OSInfo) []provider.Provider {
	var providers []provider.Provider
	
	if osInfo.IsMacOS() {
		p := provider.NewBrewProvider()
		if p.IsAvailable() {
			providers = append(providers, p)
		}
	}
	
	if osInfo.IsWindows() {
		p := provider.NewWinGetProvider()
		if p.IsAvailable() {
			providers = append(providers, p)
		}
	}
	
	if osInfo.IsLinux() {
		apt := provider.NewAptProvider()
		if apt.IsAvailable() {
			providers = append(providers, apt)
		}
		
		snap := provider.NewSnapProvider()
		if snap.IsAvailable() {
			providers = append(providers, snap)
		}
	}
	
	return providers
}
