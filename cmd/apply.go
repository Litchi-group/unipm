package cmd

import (
	"fmt"

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
	// TODO: Implement apply logic
	if dryRun {
		fmt.Println("Dry run mode enabled. Nothing will be executed.")
		fmt.Println()
	}
	
	fmt.Println("Apply not yet implemented.")
	fmt.Println()
	fmt.Println("Future output:")
	fmt.Println("  Installing vscode...")
	fmt.Println("    ✓ brew install --cask visual-studio-code")
	fmt.Println("  Installing git...")
	fmt.Println("    ⊙ Already installed")
	fmt.Println("  Installing node...")
	fmt.Println("    ✓ brew install node")
	fmt.Println()
	fmt.Println("Done! 2 installed, 1 skipped.")
	
	return nil
}
