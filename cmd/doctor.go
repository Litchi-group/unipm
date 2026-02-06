package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system requirements and package manager availability",
	Long: `Verifies that required package managers are installed and available
on the current system. Provides actionable guidance if tools are missing.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDoctor()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

func runDoctor() error {
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Println()
	
	// TODO: Implement actual provider checks
	fmt.Println("Package manager checks not yet implemented.")
	fmt.Println()
	fmt.Println("Future output:")
	
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("  ✓ Homebrew: installed (/opt/homebrew/bin/brew)")
	case "windows":
		fmt.Println("  ✓ WinGet: installed")
	case "linux":
		fmt.Println("  ✓ apt: installed (/usr/bin/apt)")
		fmt.Println("  ✓ snap: installed (/usr/bin/snap)")
	}
	
	fmt.Println()
	fmt.Println("All required tools are available.")
	
	return nil
}
