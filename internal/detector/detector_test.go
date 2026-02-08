package detector

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectOS(t *testing.T) {
	osInfo := DetectOS()
	
	assert.NotNil(t, osInfo)
	assert.NotEmpty(t, osInfo.Platform)
	assert.NotEmpty(t, osInfo.Arch)
	
	// Verify platform is one of the supported values
	assert.Contains(t, []string{"darwin", "windows", "linux"}, osInfo.Platform)
}

func TestOSInfo_IsMacOS(t *testing.T) {
	tests := []struct {
		name     string
		osInfo   *OSInfo
		expected bool
	}{
		{
			name:     "macOS",
			osInfo:   &OSInfo{Platform: "darwin"},
			expected: true,
		},
		{
			name:     "Windows",
			osInfo:   &OSInfo{Platform: "windows"},
			expected: false,
		},
		{
			name:     "Linux",
			osInfo:   &OSInfo{Platform: "linux"},
			expected: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.osInfo.IsMacOS()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOSInfo_IsWindows(t *testing.T) {
	tests := []struct {
		name     string
		osInfo   *OSInfo
		expected bool
	}{
		{
			name:     "macOS",
			osInfo:   &OSInfo{Platform: "darwin"},
			expected: false,
		},
		{
			name:     "Windows",
			osInfo:   &OSInfo{Platform: "windows"},
			expected: true,
		},
		{
			name:     "Linux",
			osInfo:   &OSInfo{Platform: "linux"},
			expected: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.osInfo.IsWindows()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOSInfo_IsLinux(t *testing.T) {
	tests := []struct {
		name     string
		osInfo   *OSInfo
		expected bool
	}{
		{
			name:     "macOS",
			osInfo:   &OSInfo{Platform: "darwin"},
			expected: false,
		},
		{
			name:     "Windows",
			osInfo:   &OSInfo{Platform: "windows"},
			expected: false,
		},
		{
			name:     "Linux",
			osInfo:   &OSInfo{Platform: "linux"},
			expected: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.osInfo.IsLinux()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOSInfo_String(t *testing.T) {
	osInfo := &OSInfo{
		Platform: "darwin",
		Arch:     "arm64",
	}
	
	result := osInfo.String()
	assert.Equal(t, "darwin", result)
	
	// Linux with distro
	osInfo2 := &OSInfo{
		Platform: "linux",
		Distro:   "ubuntu",
	}
	
	result2 := osInfo2.String()
	assert.Equal(t, "linux (ubuntu)", result2)
}

func TestDetectOS_RealSystem(t *testing.T) {
	osInfo := DetectOS()
	
	// Should match runtime.GOOS
	assert.Equal(t, runtime.GOOS, osInfo.Platform)
	
	// Should match runtime.GOARCH
	assert.Equal(t, runtime.GOARCH, osInfo.Arch)
}

func TestDetectLinuxDistro(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Linux-specific test")
	}
	
	osInfo := DetectOS()
	
	// On Linux, Distro should be populated
	assert.NotEmpty(t, osInfo.Distro)
}

func TestOSInfo_Equality(t *testing.T) {
	os1 := &OSInfo{Platform: "darwin", Arch: "arm64"}
	os2 := &OSInfo{Platform: "darwin", Arch: "arm64"}
	os3 := &OSInfo{Platform: "linux", Arch: "amd64"}
	
	assert.Equal(t, os1.Platform, os2.Platform)
	assert.NotEqual(t, os1.Platform, os3.Platform)
}
