package cmd

import (
	"fmt"
	goerrors "errors"

	"github.com/Litchi-group/unipm/internal/errors"
)

// handleError provides user-friendly error messages based on error type
func handleError(err error) error {
	if err == nil {
		return nil
	}
	
	// Check for specific error types
	var notFoundErr *errors.NotFoundError
	if goerrors.As(err, &notFoundErr) {
		return fmt.Errorf("❌ Package not found: %s\n\nTry searching available packages with: unipm search %s", notFoundErr.PackageID, notFoundErr.PackageID)
	}
	
	var circularErr *errors.CircularDependencyError
	if goerrors.As(err, &circularErr) {
		return fmt.Errorf("❌ Circular dependency detected!\n\nDependency cycle: %s\n\nPlease remove one of these dependencies or contact the package maintainer.", circularErr.Error())
	}
	
	var providerErr *errors.ProviderUnavailableError
	if goerrors.As(err, &providerErr) {
		return fmt.Errorf("❌ Provider not available: %s\n\n%s\n\nRun 'unipm doctor' to check your system setup.", providerErr.ProviderName, providerErr.Message)
	}
	
	var depErr *errors.DependencyError
	if goerrors.As(err, &depErr) {
		return fmt.Errorf("❌ Dependency error for '%s': %s", depErr.PackageID, depErr.Message)
	}
	
	var networkErr *errors.NetworkError
	if goerrors.As(err, &networkErr) {
		return fmt.Errorf("❌ Network error accessing registry:\n\nURL: %s\nError: %s\n\nPlease check your internet connection.", networkErr.URL, networkErr.Message)
	}
	
	var configErr *errors.ConfigError
	if goerrors.As(err, &configErr) {
		return fmt.Errorf("❌ Configuration error in '%s': %s", configErr.FilePath, configErr.Message)
	}
	
	var installErr *errors.InstallError
	if goerrors.As(err, &installErr) {
		return fmt.Errorf("❌ Failed to install '%s' via %s:\n\n%s", installErr.PackageID, installErr.Provider, installErr.Message)
	}
	
	// Default error
	return err
}
