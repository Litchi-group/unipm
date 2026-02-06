package detector

import (
	"os"
	"runtime"
	"strings"
)

// OSInfo contains information about the current operating system
type OSInfo struct {
	Platform string // "darwin", "windows", "linux"
	Distro   string // "ubuntu", "debian", etc. (Linux only)
	Arch     string // "amd64", "arm64", etc.
}

// DetectOS detects the current operating system
func DetectOS() *OSInfo {
	info := &OSInfo{
		Platform: runtime.GOOS,
		Arch:     runtime.GOARCH,
	}

	// Detect Linux distribution if applicable
	if info.Platform == "linux" {
		info.Distro = detectLinuxDistro()
	}

	return info
}

// detectLinuxDistro attempts to detect the Linux distribution
func detectLinuxDistro() string {
	// Try /etc/os-release (most modern distros)
	if data, err := os.ReadFile("/etc/os-release"); err == nil {
		content := string(data)
		
		// Look for ID= line
		for _, line := range strings.Split(content, "\n") {
			if strings.HasPrefix(line, "ID=") {
				distro := strings.TrimPrefix(line, "ID=")
				distro = strings.Trim(distro, `"`)
				return distro
			}
		}
	}

	// Try /etc/lsb-release (older Ubuntu/Debian)
	if data, err := os.ReadFile("/etc/lsb-release"); err == nil {
		content := string(data)
		
		for _, line := range strings.Split(content, "\n") {
			if strings.HasPrefix(line, "DISTRIB_ID=") {
				distro := strings.TrimPrefix(line, "DISTRIB_ID=")
				distro = strings.ToLower(distro)
				return distro
			}
		}
	}

	return "unknown"
}

// IsMacOS returns true if running on macOS
func (o *OSInfo) IsMacOS() bool {
	return o.Platform == "darwin"
}

// IsWindows returns true if running on Windows
func (o *OSInfo) IsWindows() bool {
	return o.Platform == "windows"
}

// IsLinux returns true if running on Linux
func (o *OSInfo) IsLinux() bool {
	return o.Platform == "linux"
}

// IsUbuntu returns true if running on Ubuntu
func (o *OSInfo) IsUbuntu() bool {
	return o.IsLinux() && strings.Contains(strings.ToLower(o.Distro), "ubuntu")
}

// IsDebian returns true if running on Debian
func (o *OSInfo) IsDebian() bool {
	return o.IsLinux() && o.Distro == "debian"
}

// String returns a human-readable string representation
func (o *OSInfo) String() string {
	if o.IsLinux() && o.Distro != "" && o.Distro != "unknown" {
		return o.Platform + " (" + o.Distro + ")"
	}
	return o.Platform
}
