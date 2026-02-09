package cmd

import (
	"fmt"
	"strings"

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

	fmt.Println("🏥 unipm System Check")
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Println()

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

	fmt.Println("Package Managers:")
	fmt.Println("-" + strings.Repeat("-", 50))
	fmt.Println()

	// Check each provider
	allAvailable := true
	missingProviders := []provider.Provider{}
	availableCount := 0

	for _, p := range providers {
		if p.IsAvailable() {
			fmt.Printf("✅ %s: available\n", p.Name())
			availableCount++
		} else {
			fmt.Printf("❌ %s: not found\n", p.Name())
			allAvailable = false
			missingProviders = append(missingProviders, p)
		}
	}

	fmt.Println()
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Println()

	if allAvailable {
		fmt.Println("✅ All required tools are available!")
		fmt.Println()
		fmt.Println("You're ready to use unipm:")
		fmt.Println("  • Run 'unipm init' to create a devpack.yaml")
		fmt.Println("  • Run 'unipm search <package>' to find packages")
		fmt.Println("  • Run 'unipm --help' for more commands")
	} else {
		// For Linux, if at least one package manager is available, it's OK
		if osInfo.IsLinux() && availableCount > 0 {
			fmt.Println("✅ System check passed!")
			fmt.Println()
			fmt.Printf("You have %d/%d package managers available.\n", availableCount, len(providers))
			fmt.Println("This is sufficient to use unipm.")
			if len(missingProviders) > 0 {
				fmt.Println()
				fmt.Println("Optional: Install missing package managers for more coverage:")
				for _, p := range missingProviders {
					fmt.Println()
					fmt.Println(provider.GetInstallationGuide(p.Name()))
				}
			}
		} else {
			fmt.Println("❌ System check failed!")
			fmt.Println()
			fmt.Println("Missing required package managers:")
			fmt.Println()

			// Show installation guides for missing tools
			for _, p := range missingProviders {
				fmt.Println(provider.GetInstallationGuide(p.Name()))
				fmt.Println()
			}

			fmt.Println("After installing the required tools, run 'unipm doctor' again.")
			return fmt.Errorf("missing required package managers")
		}
	}

	return nil
}
