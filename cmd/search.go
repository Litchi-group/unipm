package cmd

import (
	"fmt"
	"strings"

	"github.com/Litchi-group/unipm/internal/registry"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for packages in the registry",
	Long:  `Searches the package registry for packages matching the query.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.Join(args, " ")
		return runSearch(query)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func runSearch(query string) error {
	reg := registry.NewRegistry()

	// Load package index
	packages, err := reg.LoadIndex()
	if err != nil {
		return fmt.Errorf("failed to load package index: %w", err)
	}

	// Filter packages by query
	query = strings.ToLower(query)
	matches := []registry.PackageInfo{}

	for _, pkg := range packages {
		if strings.Contains(strings.ToLower(pkg.ID), query) ||
			strings.Contains(strings.ToLower(pkg.Name), query) {
			matches = append(matches, pkg)
		}
	}

	if len(matches) == 0 {
		fmt.Printf("No packages found matching '%s'\n", query)
		return nil
	}

	fmt.Printf("Found %d package(s):\n\n", len(matches))

	for _, pkg := range matches {
		fmt.Printf("  %s - %s\n", pkg.ID, pkg.Name)
	}

	fmt.Printf("\nUse 'unipm info <package>' for more details.\n")

	return nil
}
