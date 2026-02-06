package planner

import (
	"fmt"

	"github.com/Litchi-group/unipm/internal/detector"
	"github.com/Litchi-group/unipm/internal/provider"
	"github.com/Litchi-group/unipm/internal/registry"
)

// InstallTask represents a single installation task
type InstallTask struct {
	PackageID string
	Spec      *provider.ProviderSpec
	Provider  provider.Provider
	Installed bool
}

// Plan represents an installation plan
type Plan struct {
	Tasks  []*InstallTask
	OSInfo *detector.OSInfo
}

// Planner generates installation plans
type Planner struct {
	registry *registry.Registry
	resolver *registry.Resolver
	osInfo   *detector.OSInfo
}

// NewPlanner creates a new Planner
func NewPlanner(reg *registry.Registry, osInfo *detector.OSInfo) *Planner {
	return &Planner{
		registry: reg,
		resolver: registry.NewResolver(reg, osInfo),
		osInfo:   osInfo,
	}
}

// CreatePlan creates an installation plan for the given package IDs
func (p *Planner) CreatePlan(packageIDs []string) (*Plan, error) {
	plan := &Plan{
		Tasks:  make([]*InstallTask, 0, len(packageIDs)),
		OSInfo: p.osInfo,
	}
	
	for _, packageID := range packageIDs {
		// Resolve package to provider spec
		spec, err := p.resolver.Resolve(packageID)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve %s: %w", packageID, err)
		}
		
		// Get provider instance
		prov, err := provider.GetProviderByType(spec.Type)
		if err != nil {
			return nil, fmt.Errorf("failed to get provider for %s: %w", packageID, err)
		}
		
		// Check if provider is available
		if !prov.IsAvailable() {
			return nil, fmt.Errorf("provider %s is not available for %s", prov.Name(), packageID)
		}
		
		// Check if already installed
		installed := prov.IsInstalled(*spec)
		
		task := &InstallTask{
			PackageID: packageID,
			Spec:      spec,
			Provider:  prov,
			Installed: installed,
		}
		
		plan.Tasks = append(plan.Tasks, task)
	}
	
	return plan, nil
}

// Execute executes the installation plan
func (plan *Plan) Execute(dryRun bool) error {
	installedCount := 0
	skippedCount := 0
	
	for _, task := range plan.Tasks {
		if task.Installed {
			fmt.Printf("Installing %s...\n", task.PackageID)
			fmt.Printf("  ⊙ Already installed\n")
			skippedCount++
			continue
		}
		
		fmt.Printf("Installing %s...\n", task.PackageID)
		
		if dryRun {
			fmt.Printf("  [dry-run] %s\n", task.Provider.InstallCommand(*task.Spec))
			installedCount++
			continue
		}
		
		// Execute installation
		if err := task.Provider.Install(*task.Spec); err != nil {
			return fmt.Errorf("failed to install %s: %w", task.PackageID, err)
		}
		
		fmt.Printf("  ✓ Installed\n")
		installedCount++
	}
	
	fmt.Println()
	
	if dryRun {
		fmt.Printf("Dry run complete. Would install %d, skip %d.\n", installedCount, skippedCount)
	} else {
		fmt.Printf("Done! %d installed, %d skipped.\n", installedCount, skippedCount)
	}
	
	return nil
}

// Print prints the plan without executing
func (plan *Plan) Print() {
	fmt.Printf("Plan for %s:\n\n", plan.OSInfo.String())
	
	for _, task := range plan.Tasks {
		cmd := task.Provider.InstallCommand(*task.Spec)
		status := ""
		if task.Installed {
			status = " (already installed)"
		}
		fmt.Printf("  %s → %s%s\n", task.PackageID, cmd, status)
	}
	
	fmt.Println()
	fmt.Println("To apply this plan, run 'unipm apply'.")
}
