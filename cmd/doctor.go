package cmd

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/provider"
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
	// Detect OS
	osInfo := detector.DetectOS()

	fmt.Printf("OS: %s\n", osInfo.String())
	fmt.Printf("Architecture: %s\n", osInfo.Arch)
	fmt.Println()

	// Check providers based on OS
	var providers []provider.Provider

	switch {
	case osInfo.IsMacOS():
		providers = []provider.Provider{
			provider.NewBrewProvider(),
		}
	case osInfo.IsWindows():
		providers = []provider.Provider{
			provider.NewWinGetProvider(),
		}
	case osInfo.IsLinux():
		providers = []provider.Provider{
			provider.NewAptProvider(),
			provider.NewSnapProvider(),
		}
	}

	// Check each provider
	allAvailable := true
	for _, p := range providers {
		if p.IsAvailable() {
			fmt.Printf("✓ %s: available\n", p.Name())
		} else {
			fmt.Printf("✗ %s: not found\n", p.Name())
			allAvailable = false
		}
	}

	fmt.Println()

	if allAvailable {
		fmt.Println("All required tools are available.")
	} else {
		fmt.Println("Some required tools are missing.")
		fmt.Println()

		// Show installation guides for missing tools
		for _, p := range providers {
			if !p.IsAvailable() {
				fmt.Println(provider.GetInstallationGuide(p.Name()))
				fmt.Println()
			}
		}
	}

	return nil
}
