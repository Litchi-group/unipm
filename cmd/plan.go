package cmd

import (
	"fmt"

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
	// TODO: Implement plan generation
	fmt.Println("Plan generation not yet implemented.")
	fmt.Println()
	fmt.Println("Future output:")
	fmt.Println("  vscode → brew install --cask visual-studio-code")
	fmt.Println("  git    → brew install git")
	fmt.Println("  node   → brew install node")
	fmt.Println()
	fmt.Println("To apply this plan, run 'unipm apply'.")
	
	return nil
}
