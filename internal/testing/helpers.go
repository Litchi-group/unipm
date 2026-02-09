package testing

import (
	"os"
	"path/filepath"
	"testing"
)

// TempDir creates a temporary directory for testing
// Returns the path and a cleanup function
func TempDir(t *testing.T) (string, func()) {
	t.Helper()

	dir, err := os.MkdirTemp("", "unipm-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	cleanup := func() {
		_ = os.RemoveAll(dir)
	}

	return dir, cleanup
}

// WriteFile writes content to a file in the temp directory
func WriteFile(t *testing.T, dir, filename, content string) string {
	t.Helper()

	path := filepath.Join(dir, filename)

	// Create parent directories if needed
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write file %s: %v", path, err)
	}

	return path
}

// SetEnv sets an environment variable for the duration of the test
func SetEnv(t *testing.T, key, value string) func() {
	t.Helper()

	old := os.Getenv(key)
	_ = os.Setenv(key, value)

	return func() {
		if old == "" {
			_ = os.Unsetenv(key)
		} else {
			_ = os.Setenv(key, old)
		}
	}
}

// CaptureStdout captures stdout during a function call
// Returns the captured output
func CaptureStdout(t *testing.T, fn func()) string {
	t.Helper()

	// Create a pipe
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	// Save the original stdout
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	// Redirect stdout to the write end of the pipe
	os.Stdout = w

	// Run the function
	fn()

	// Close the write end
	_ = w.Close()

	// Read the captured output
	buf := make([]byte, 4096)
	n, _ := r.Read(buf)

	return string(buf[:n])
}
