package registry

// RegistryInterface defines the interface for package registries
type RegistryInterface interface {
	// LoadPackage loads a package definition by ID
	LoadPackage(packageID string) (*Package, error)

	// LoadIndex loads the package index
	LoadIndex() ([]PackageInfo, error)
}

// Ensure Registry implements RegistryInterface
var _ RegistryInterface = (*Registry)(nil)
