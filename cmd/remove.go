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

var removeYes bool

var removeCmd = &cobra.Command{
	Use:   "remove <package...>",
	Short: "Remove packages",
	Long: `Removes specified packages using the native package manager.
Prompts for confirmation before removing unless --yes is specified.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRemove(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&removeYes, "yes", "y", false, "Skip confirmation prompt")
}

func runRemove(packageIDs []string) error {
	// Load devpack.yaml to verify packages
	devpack, err := config.Load("devpack.yaml")
	if err != nil {
		return fmt.Errorf("failed to load devpack.yaml: %w", err)
	}

	// Check if packages are in devpack.yaml
	for _, pkg := range packageIDs {
		found := false
		for _, devPkg := range devpack.Apps {
			if devPkg == pkg {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Warning: %s is not in devpack.yaml\n", pkg)
		}
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

	fmt.Printf("Packages to remove:\n\n")

	toRemove := 0
	for _, task := range plan.Tasks {
		if task.Installed {
			fmt.Printf("  %s → %s\n", task.PackageID, task.Provider.RemoveCommand(*task.Spec))
			toRemove++
		} else {
			fmt.Printf("  %s (not installed)\n", task.PackageID)
		}
	}

	fmt.Println()

	if toRemove == 0 {
		fmt.Println("No packages to remove.")
		return nil
	}

	// Prompt for confirmation unless --yes
	if !removeYes {
		fmt.Printf("Do you want to proceed with removing %d package(s)? [y/N]: ", toRemove)

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

	// Execute removals
	removedCount := 0
	notInstalledCount := 0

	for _, task := range plan.Tasks {
		if !task.Installed {
			fmt.Printf("Removing %s...\n", task.PackageID)
			fmt.Printf("  ⊙ Not installed\n")
			notInstalledCount++
			continue
		}

		fmt.Printf("Removing %s...\n", task.PackageID)

		if err := task.Provider.Remove(*task.Spec); err != nil {
			return fmt.Errorf("failed to remove %s: %w", task.PackageID, err)
		}

		fmt.Printf("  ✓ Removed\n")
		removedCount++
	}

	fmt.Println()
	fmt.Printf("Done! %d removed, %d not installed.\n", removedCount, notInstalledCount)

	return nil
}
