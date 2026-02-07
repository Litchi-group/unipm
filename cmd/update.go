package cmd

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/config"
	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/planner"
	"github.com/Litchi-group/unipm/internal/registry"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [package...]",
	Short: "Update packages",
	Long: `Updates all packages in devpack.yaml, or specific packages if provided.
Uses the native package manager's update command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUpdate(args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(packageIDs []string) error {
	// Load devpack.yaml
	devpack, err := config.Load("devpack.yaml")
	if err != nil {
		return fmt.Errorf("failed to load devpack.yaml: %w", err)
	}
	
	// If no packages specified, update all
	if len(packageIDs) == 0 {
		packageIDs = devpack.Apps
	}
	
	if len(packageIDs) == 0 {
		fmt.Println("No packages to update")
		return nil
	}
	
	// Detect OS
	osInfo := detector.DetectOS()
	
	// Create planner
	reg := registry.NewRegistry()
	plnr := planner.NewPlanner(reg, osInfo)
	
	// Create plan
	plan, err := plnr.CreatePlan(packageIDs)
	if err != nil {
		return err
	}
	
	fmt.Printf("Updating %d package(s)...\n\n", len(plan.Tasks))
	
	updatedCount := 0
	notInstalledCount := 0
	
	for _, task := range plan.Tasks {
		if !task.Installed {
			fmt.Printf("Updating %s...\n", task.PackageID)
			fmt.Printf("  ⊙ Not installed (use 'unipm apply' to install)\n")
			notInstalledCount++
			continue
		}
		
		fmt.Printf("Updating %s...\n", task.PackageID)
		
		// Get update command (same as install for most package managers)
		// They handle updates when package is already installed
		if err := task.Provider.Install(*task.Spec); err != nil {
			return fmt.Errorf("failed to update %s: %w", task.PackageID, err)
		}
		
		fmt.Printf("  ✓ Updated\n")
		updatedCount++
	}
	
	fmt.Println()
	fmt.Printf("Done! %d updated, %d not installed.\n", updatedCount, notInstalledCount)
	
	return nil
}
