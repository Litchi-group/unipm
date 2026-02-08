package registry

import "fmt"

// MockRegistry is a mock implementation of Registry for testing
type MockRegistry struct {
	Packages      map[string]*Package
	IndexPackages []PackageInfo
	LoadError     error
	IndexError    error
}

// NewMockRegistry creates a new mock registry
func NewMockRegistry() *MockRegistry {
	return &MockRegistry{
		Packages:      make(map[string]*Package),
		IndexPackages: []PackageInfo{},
	}
}

// LoadPackage returns a pre-configured package or an error
func (m *MockRegistry) LoadPackage(packageID string) (*Package, error) {
	if m.LoadError != nil {
		return nil, m.LoadError
	}
	
	pkg, ok := m.Packages[packageID]
	if !ok {
		return nil, fmt.Errorf("package %s not found", packageID)
	}
	
	return pkg, nil
}

// LoadIndex returns a pre-configured index or an error
func (m *MockRegistry) LoadIndex() ([]PackageInfo, error) {
	if m.IndexError != nil {
		return nil, m.IndexError
	}
	
	return m.IndexPackages, nil
}

// AddPackage adds a package to the mock registry for testing
func (m *MockRegistry) AddPackage(pkg *Package) {
	m.Packages[pkg.ID] = pkg
}

// AddIndexPackage adds a package info to the mock index
func (m *MockRegistry) AddIndexPackage(info PackageInfo) {
	m.IndexPackages = append(m.IndexPackages, info)
}

// SetLoadError sets the error to return from LoadPackage
func (m *MockRegistry) SetLoadError(err error) {
	m.LoadError = err
}

// SetIndexError sets the error to return from LoadIndex
func (m *MockRegistry) SetIndexError(err error) {
	m.IndexError = err
}
