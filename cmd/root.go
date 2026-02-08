package cmd

import (
	"fmt"
	"os"

	"github.com/Litchi-group/unipm/internal/logger"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "unipm",
	Short: "Universal package manager orchestrator",
	Long: `unipm - Universal package manager orchestrator for Windows, macOS, and Linux.

unipm orchestrates native package managers (Homebrew, WinGet, apt, snap) 
to provide a unified, declarative interface for cross-platform dev environment setup.

Write once. Set up anywhere.`,
	Version: "0.1.1",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logger.SetLevel(logger.LevelDebug)
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
}
