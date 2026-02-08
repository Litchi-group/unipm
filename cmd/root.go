package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unipm",
	Short: "Universal package manager orchestrator",
	Long: `unipm - Universal package manager orchestrator for Windows, macOS, and Linux.

unipm orchestrates native package managers (Homebrew, WinGet, apt, snap) 
to provide a unified, declarative interface for cross-platform dev environment setup.

Write once. Set up anywhere.`,
	Version: "0.1.2",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.unipm.yaml)")
}
