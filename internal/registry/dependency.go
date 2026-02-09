package registry

import (
	"github.com/Litchi-group/unipm/internal/errors"
)

// DependencyResolver resolves package dependencies and returns installation order
type DependencyResolver struct {
	registry *Registry
}

// NewDependencyResolver creates a new dependency resolver
func NewDependencyResolver(registry *Registry) *DependencyResolver {
	return &DependencyResolver{
		registry: registry,
	}
}

// Resolve resolves dependencies and returns packages in installation order
// Uses topological sort to handle dependency chains
func (dr *DependencyResolver) Resolve(packageIDs []string) ([]string, error) {
	visited := make(map[string]bool)
	visiting := make(map[string]bool)
	path := []string{}
	result := []string{}

	var visit func(string) error
	visit = func(pkgID string) error {
		if visited[pkgID] {
			return nil
		}

		if visiting[pkgID] {
			// Found a cycle - build the cycle path
			cycleStart := -1
			for i, p := range path {
				if p == pkgID {
					cycleStart = i
					break
				}
			}

			var cycle []string
			if cycleStart >= 0 {
				cycle = append(cycle, path[cycleStart:]...)
				cycle = append(cycle, pkgID) // Close the cycle
			} else {
				cycle = []string{pkgID, pkgID}
			}

			return errors.NewCircularDependencyError(cycle)
		}

		visiting[pkgID] = true
		path = append(path, pkgID)

		// Load package to check dependencies
		pkg, err := dr.registry.LoadPackage(pkgID)
		if err != nil {
			return errors.NewDependencyError(pkgID, "failed to load package", err)
		}

		// Visit dependencies first
		for _, depID := range pkg.Dependencies {
			if err := visit(depID); err != nil {
				return err
			}
		}

		path = path[:len(path)-1]
		visiting[pkgID] = false
		visited[pkgID] = true
		result = append(result, pkgID)

		return nil
	}

	// Visit all requested packages
	for _, pkgID := range packageIDs {
		if err := visit(pkgID); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetDependencyTree returns a map of package -> dependencies for visualization
func (dr *DependencyResolver) GetDependencyTree(packageIDs []string) (map[string][]string, error) {
	tree := make(map[string][]string)
	visited := make(map[string]bool)

	var visit func(string) error
	visit = func(pkgID string) error {
		if visited[pkgID] {
			return nil
		}
		visited[pkgID] = true

		pkg, err := dr.registry.LoadPackage(pkgID)
		if err != nil {
			return err
		}

		tree[pkgID] = pkg.Dependencies

		for _, depID := range pkg.Dependencies {
			if err := visit(depID); err != nil {
				return err
			}
		}

		return nil
	}

	for _, pkgID := range packageIDs {
		if err := visit(pkgID); err != nil {
			return nil, err
		}
	}

	return tree, nil
}
