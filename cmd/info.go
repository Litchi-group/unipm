package cmd

import (
	"fmt"
	"strings"

	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/registry"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <package>",
	Short: "Show detailed information about a package",
	Long:  `Displays detailed information about a package including available providers and dependencies.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInfo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func runInfo(packageID string) error {
	reg := registry.NewRegistry()
	
	// Load package
	pkg, err := reg.LoadPackage(packageID)
	if err != nil {
		return fmt.Errorf("failed to load package: %w", err)
	}
	
	// Display information
	fmt.Printf("Package: %s\n", pkg.Name)
	fmt.Printf("ID: %s\n", pkg.ID)
	
	if pkg.Homepage != "" {
		fmt.Printf("Homepage: %s\n", pkg.Homepage)
	}
	
	// Dependencies
	if len(pkg.Dependencies) > 0 {
		fmt.Printf("\nDependencies:\n")
		for _, dep := range pkg.Dependencies {
			fmt.Printf("  - %s\n", dep)
		}
	}
	
	// Providers
	fmt.Printf("\nProviders:\n")
	
	osInfo := detector.DetectOS()
	
	for osKey, providers := range pkg.Providers {
		fmt.Printf("  %s:\n", osKey)
		for _, p := range providers {
			fmt.Printf("    - %s: %s\n", p.Type, p.Name)
			
			// Highlight if this is the current OS
			if strings.Contains(osInfo.String(), osKey) {
				fmt.Printf("      (available on your system)\n")
			}
		}
	}
	
	return nil
}
