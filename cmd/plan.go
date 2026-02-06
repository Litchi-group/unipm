package cmd

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/config"
	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/provider"
	"github.com/Litchi-group/unipm/internal/registry"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Generate an installation plan from devpack.yaml",
	Long: `Reads devpack.yaml and generates a detailed installation plan
showing what would be installed on the current OS.

This is a non-destructive operation that only displays the plan.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runPlan()
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}

func runPlan() error {
	// Load devpack.yaml
	devpack, err := config.Load("devpack.yaml")
	if err != nil {
		return fmt.Errorf("failed to load devpack.yaml: %w", err)
	}
	
	if len(devpack.Apps) == 0 {
		fmt.Println("No packages specified in devpack.yaml")
		return nil
	}
	
	// Detect OS
	osInfo := detector.DetectOS()
	
	// Create registry and resolver
	reg := registry.NewRegistry()
	resolver := registry.NewResolver(reg, osInfo)
	
	// Resolve packages
	fmt.Printf("Plan for %s:\n\n", osInfo.String())
	
	for _, packageID := range devpack.Apps {
		spec, err := resolver.Resolve(packageID)
		if err != nil {
			fmt.Printf("  ✗ %s: %v\n", packageID, err)
			continue
		}
		
		// Get provider
		prov, err := provider.GetProviderByType(spec.Type)
		if err != nil {
			fmt.Printf("  ✗ %s: %v\n", packageID, err)
			continue
		}
		
		// Format command
		cmd := prov.InstallCommand(*spec)
		fmt.Printf("  %s → %s\n", packageID, cmd)
	}
	
	fmt.Println()
	fmt.Println("To apply this plan, run 'unipm apply'.")
	
	return nil
}
