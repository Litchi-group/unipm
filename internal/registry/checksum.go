package registry

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Litchi-group/unipm/internal/logger"
	"gopkg.in/yaml.v3"
)

// VerifyChecksum verifies the checksum of package data
func VerifyChecksum(data []byte, pkg *Package) bool {
	// If no checksum is provided, skip verification
	if pkg.Checksum == "" {
		logger.Debug("No checksum provided for package %s, skipping verification", pkg.ID)
		return true
	}
	
	// Calculate checksum of the data (excluding the checksum field itself)
	calculatedChecksum := calculatePackageChecksum(data)
	
	// Compare checksums
	if calculatedChecksum != pkg.Checksum {
		logger.Warn("Checksum mismatch for package %s: expected %s, got %s", 
			pkg.ID, pkg.Checksum, calculatedChecksum)
		return false
	}
	
	logger.Debug("Checksum verified for package %s", pkg.ID)
	return true
}

// calculatePackageChecksum calculates the SHA256 checksum of package data
func calculatePackageChecksum(data []byte) string {
	// Parse the YAML to remove the checksum field
	var pkg Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		// If parsing fails, calculate checksum of raw data
		return calculateRawChecksum(data)
	}
	
	// Remove checksum field
	pkg.Checksum = ""
	
	// Marshal back to YAML
	cleanData, err := yaml.Marshal(&pkg)
	if err != nil {
		return calculateRawChecksum(data)
	}
	
	return calculateRawChecksum(cleanData)
}

// calculateRawChecksum calculates SHA256 checksum of raw data
func calculateRawChecksum(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// GenerateChecksum generates a checksum for a package file
func GenerateChecksum(data []byte) string {
	return calculatePackageChecksum(data)
}
