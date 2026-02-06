package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new devpack.yaml file",
	Long: `Creates a devpack.yaml configuration file in the current directory
with a minimal example configuration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit() error {
	const devpackTemplate = `apps:
  - vscode
  - git
  - node
`

	filename := "devpack.yaml"
	
	// Check if file already exists
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("%s already exists", filename)
	}

	// Write template to file
	if err := os.WriteFile(filename, []byte(devpackTemplate), 0644); err != nil {
		return fmt.Errorf("failed to create %s: %w", filename, err)
	}

	fmt.Printf("✓ Created %s\n", filename)
	fmt.Println("Edit the file and run 'unipm plan' to preview.")
	
	return nil
}
