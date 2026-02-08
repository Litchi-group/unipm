package cmd

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/config"
	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/planner"
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
	
	// Print plan
	plan.Print()
	
	return nil
}
