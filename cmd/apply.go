package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Litchi-group/unipm/internal/config"
	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/planner"
	"github.com/Litchi-group/unipm/internal/registry"
	"github.com/spf13/cobra"
)

var (
	dryRun bool
	yes    bool
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the installation plan",
	Long: `Executes the installation plan by invoking native package managers.
Skips packages that are already installed.

By default, prompts for confirmation before executing.
Use --yes to skip confirmation.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApply()
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	
	applyCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be done without executing")
	applyCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip confirmation prompt")
}

func runApply() error {
	// Load devpack.yaml
	devpack, err := config.Load("devpack.yaml")
	if err != nil {
		return handleError(fmt.Errorf("failed to load devpack.yaml: %w", err))
	}
	
	if len(devpack.Apps) == 0 {
		fmt.Println("No packages specified in devpack.yaml")
		return nil
	}
	
	// Detect OS
	osInfo := detector.DetectOS()
	
	// Create planner
	reg := registry.NewRegistry()
	plnr := planner.NewPlanner(reg, osInfo)
	
	// Create plan
	plan, err := plnr.CreatePlan(devpack.Apps)
	if err != nil {
		return handleError(err)
	}
	
	// Show plan summary
	fmt.Printf("Plan for %s:\n\n", osInfo.String())
	
	newInstalls := 0
	for _, task := range plan.Tasks {
		if !task.Installed {
			newInstalls++
			fmt.Printf("  %s → %s\n", task.PackageID, task.Provider.InstallCommand(*task.Spec))
		} else {
			fmt.Printf("  %s (already installed)\n", task.PackageID)
		}
	}
	
	fmt.Println()
	
	if newInstalls == 0 {
		fmt.Println("All packages are already installed.")
		return nil
	}
	
	// Prompt for confirmation unless --yes or --dry-run
	if !dryRun && !yes {
		fmt.Printf("Do you want to proceed with installing %d package(s)? [y/N]: ", newInstalls)
		
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" && response != "yes" {
			fmt.Println("Cancelled.")
			return nil
		}
		
		fmt.Println()
	}
	
	// Execute plan
	if dryRun {
		fmt.Println("Dry run mode enabled. Nothing will be executed.")
		fmt.Println()
	}
	
	return plan.Execute(dryRun)
}
