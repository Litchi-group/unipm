package cmd

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List packages in devpack.yaml",
	Long:  `Shows all packages defined in your devpack.yaml file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList() error {
	// Load devpack.yaml
	devpack, err := config.Load("devpack.yaml")
	if err != nil {
		return fmt.Errorf("failed to load devpack.yaml: %w", err)
	}
	
	if len(devpack.Apps) == 0 {
		fmt.Println("No packages defined in devpack.yaml")
		return nil
	}
	
	fmt.Printf("Packages in devpack.yaml (%d):\n\n", len(devpack.Apps))
	
	for i, pkg := range devpack.Apps {
		fmt.Printf("  %d. %s\n", i+1, pkg)
	}
	
	return nil
}
