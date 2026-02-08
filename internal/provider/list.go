package provider

import "strings"

// ListInstalled implementation for BrewProvider
func (p *BrewProvider) ListInstalled() ([]string, error) {
	output, err := execCommand("brew", "list", "--formula")
	if err != nil {
		return nil, err
	}
	
	return parseLines(output), nil
}

// ListInstalled implementation for WinGetProvider
func (p *WinGetProvider) ListInstalled() ([]string, error) {
	output, err := execCommand("winget", "list")
	if err != nil {
		return nil, err
	}
	
	// Parse winget output (skip header lines)
	var packages []string
	lines := strings.Split(output, "\n")
	
	for i, line := range lines {
		// Skip header (first 2-3 lines)
		if i < 2 {
			continue
		}
		
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "-") {
			continue
		}
		
		// Extract package name (first column)
		fields := strings.Fields(line)
		if len(fields) > 0 {
			packages = append(packages, fields[0])
		}
	}
	
	return packages, nil
}

// ListInstalled implementation for AptProvider
func (p *AptProvider) ListInstalled() ([]string, error) {
	output, err := execCommand("dpkg", "--get-selections")
	if err != nil {
		return nil, err
	}
	
	var packages []string
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, "install") {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				packages = append(packages, fields[0])
			}
		}
	}
	
	return packages, nil
}

// ListInstalled implementation for SnapProvider
func (p *SnapProvider) ListInstalled() ([]string, error) {
	output, err := execCommand("snap", "list")
	if err != nil {
		return nil, err
	}
	
	var packages []string
	lines := strings.Split(output, "\n")
	
	for i, line := range lines {
		// Skip header
		if i == 0 {
			continue
		}
		
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Extract package name (first column)
		fields := strings.Fields(line)
		if len(fields) > 0 {
			packages = append(packages, fields[0])
		}
	}
	
	return packages, nil
}

// parseLines splits output by newlines and filters empty lines
func parseLines(output string) []string {
	var result []string
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}
