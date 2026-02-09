package provider

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Litchi-group/unipm/internal/logger"
)

// BaseProvider provides common functionality for all providers
type BaseProvider struct {
	name       string
	executable string
}

// Name returns the provider name
func (p *BaseProvider) Name() string {
	return p.name
}

// IsAvailable checks if the executable is in PATH
func (p *BaseProvider) IsAvailable() bool {
	_, err := exec.LookPath(p.executable)
	return err == nil
}

// execCommand executes a command and returns the output
func execCommand(name string, args ...string) (string, error) {
	logger.Debug("Executing: %s %s", name, strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		logger.Debug("Command failed: %v, output: %s", err, string(output))
	}

	return strings.TrimSpace(string(output)), err
}

// execCommandSilent executes a command and returns only the error
func execCommandSilent(name string, args ...string) error {
	logger.Debug("Executing: %s %s", name, strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

// executeWithDisplay executes a command after displaying it
func (p *BaseProvider) executeWithDisplay(args ...string) error {
	fmt.Printf("  → %s\n", FormatCommand(p.executable, args...))
	return execCommandSilent(p.executable, args...)
}

// checkInstalled checks if a package is installed using a list command
func (p *BaseProvider) checkInstalled(args ...string) bool {
	output, err := execCommand(p.executable, args...)
	return err == nil && strings.TrimSpace(output) != ""
}

// FormatCommand formats a command for display
func FormatCommand(name string, args ...string) string {
	parts := append([]string{name}, args...)
	return strings.Join(parts, " ")
}
