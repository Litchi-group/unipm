package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Litchi-group/unipm/internal/config"
)

// loadDevpackWithPrompt loads devpack.yaml and shows a helpful message if not found
func loadDevpackWithPrompt() (*config.DevPack, error) {
	devpack, err := config.Load("devpack.yaml")
	if err != nil {
		// Check if file doesn't exist
		if os.IsNotExist(err) || strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("📦 No devpack.yaml found in current directory")
			fmt.Println("")
			fmt.Println("To get started:")
			fmt.Println("  1. Run 'unipm init' to create a new devpack.yaml")
			fmt.Println("  2. Or navigate to a directory with an existing devpack.yaml")
			fmt.Println("")
			return nil, nil // Return nil without error to exit gracefully
		}
		return nil, fmt.Errorf("failed to load devpack.yaml: %w", err)
	}
	return devpack, nil
}
