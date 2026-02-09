package cmd

import (
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Import and install packages from devpack.yaml",
	Long: `Imports a devpack.yaml file and installs the specified packages.

This is an alias for 'unipm apply' and accepts the same flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Import is just an alias for apply
		// In the future, we could add file path support
		return runApply()
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Share flags with apply
	importCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be done without executing")
	importCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip confirmation prompt")
	importCmd.Flags().StringVarP(&profile, "profile", "p", "", "Use a specific profile from devpack.yaml")
}
