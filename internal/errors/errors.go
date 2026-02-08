package errors

import (
	"fmt"
	"strings"
)

// NotFoundError indicates a package was not found
type NotFoundError struct {
	PackageID string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("package '%s' not found in registry", e.PackageID)
}

// NewNotFoundError creates a new NotFoundError
func NewNotFoundError(packageID string) *NotFoundError {
	return &NotFoundError{PackageID: packageID}
}

// ProviderUnavailableError indicates a provider is not available
type ProviderUnavailableError struct {
	ProviderName string
	Message      string
}

func (e *ProviderUnavailableError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("provider '%s' is not available: %s", e.ProviderName, e.Message)
	}
	return fmt.Sprintf("provider '%s' is not available", e.ProviderName)
}

// NewProviderUnavailableError creates a new ProviderUnavailableError
func NewProviderUnavailableError(providerName, message string) *ProviderUnavailableError {
	return &ProviderUnavailableError{
		ProviderName: providerName,
		Message:      message,
	}
}

// CircularDependencyError indicates a circular dependency was detected
type CircularDependencyError struct {
	Cycle []string
}

func (e *CircularDependencyError) Error() string {
	if len(e.Cycle) == 0 {
		return "circular dependency detected"
	}
	return fmt.Sprintf("circular dependency detected: %s", strings.Join(e.Cycle, " → "))
}

// NewCircularDependencyError creates a new CircularDependencyError
func NewCircularDependencyError(cycle []string) *CircularDependencyError {
	return &CircularDependencyError{Cycle: cycle}
}

// DependencyError indicates a dependency resolution error
type DependencyError struct {
	PackageID string
	Message   string
	Cause     error
}

func (e *DependencyError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("dependency error for '%s': %s (caused by: %v)", e.PackageID, e.Message, e.Cause)
	}
	return fmt.Sprintf("dependency error for '%s': %s", e.PackageID, e.Message)
}

func (e *DependencyError) Unwrap() error {
	return e.Cause
}

// NewDependencyError creates a new DependencyError
func NewDependencyError(packageID, message string, cause error) *DependencyError {
	return &DependencyError{
		PackageID: packageID,
		Message:   message,
		Cause:     cause,
	}
}

// ConfigError indicates a configuration error
type ConfigError struct {
	FilePath string
	Message  string
	Cause    error
}

func (e *ConfigError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("config error in '%s': %s (caused by: %v)", e.FilePath, e.Message, e.Cause)
	}
	return fmt.Sprintf("config error in '%s': %s", e.FilePath, e.Message)
}

func (e *ConfigError) Unwrap() error {
	return e.Cause
}

// NewConfigError creates a new ConfigError
func NewConfigError(filePath, message string, cause error) *ConfigError {
	return &ConfigError{
		FilePath: filePath,
		Message:  message,
		Cause:    cause,
	}
}

// NetworkError indicates a network-related error
type NetworkError struct {
	URL     string
	Message string
	Cause   error
}

func (e *NetworkError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("network error accessing '%s': %s (caused by: %v)", e.URL, e.Message, e.Cause)
	}
	return fmt.Sprintf("network error accessing '%s': %s", e.URL, e.Message)
}

func (e *NetworkError) Unwrap() error {
	return e.Cause
}

// NewNetworkError creates a new NetworkError
func NewNetworkError(url, message string, cause error) *NetworkError {
	return &NetworkError{
		URL:     url,
		Message: message,
		Cause:   cause,
	}
}

// InstallError indicates a package installation error
type InstallError struct {
	PackageID string
	Provider  string
	Message   string
	Cause     error
}

func (e *InstallError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("failed to install '%s' via %s: %s (caused by: %v)", e.PackageID, e.Provider, e.Message, e.Cause)
	}
	return fmt.Sprintf("failed to install '%s' via %s: %s", e.PackageID, e.Provider, e.Message)
}

func (e *InstallError) Unwrap() error {
	return e.Cause
}

// NewInstallError creates a new InstallError
func NewInstallError(packageID, provider, message string, cause error) *InstallError {
	return &InstallError{
		PackageID: packageID,
		Provider:  provider,
		Message:   message,
		Cause:     cause,
	}
}
