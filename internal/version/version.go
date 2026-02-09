package version

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version represents a semantic version
type Version struct {
	Major int
	Minor int
	Patch int
	Pre   string // Pre-release version
}

var versionRegex = regexp.MustCompile(`^(\d+)(?:\.(\d+))?(?:\.(\d+))?(?:-(.+))?$`)

// Parse parses a version string (e.g., "18", "18.x", "18.16.0", "1.2.3-beta")
func Parse(s string) (*Version, error) {
	s = strings.TrimSpace(s)

	// Handle wildcard versions (18.x -> 18)
	s = strings.ReplaceAll(s, ".x", "")

	matches := versionRegex.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid version string: %s", s)
	}

	v := &Version{}

	// Major version (required)
	if matches[1] != "" {
		v.Major, _ = strconv.Atoi(matches[1])
	}

	// Minor version (optional)
	if matches[2] != "" {
		v.Minor, _ = strconv.Atoi(matches[2])
	}

	// Patch version (optional)
	if matches[3] != "" {
		v.Patch, _ = strconv.Atoi(matches[3])
	}

	// Pre-release (optional)
	if matches[4] != "" {
		v.Pre = matches[4]
	}

	return v, nil
}

// String returns the string representation of the version
func (v *Version) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Pre != "" {
		s += "-" + v.Pre
	}
	return s
}

// Compare compares two versions
// Returns: -1 if v < other, 0 if v == other, 1 if v > other
func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	// Pre-release comparison
	if v.Pre == "" && other.Pre != "" {
		return 1 // Release > pre-release
	}
	if v.Pre != "" && other.Pre == "" {
		return -1 // Pre-release < release
	}
	if v.Pre != other.Pre {
		return strings.Compare(v.Pre, other.Pre)
	}

	return 0
}

// Satisfies checks if this version satisfies a constraint
// Supports: "18", "18.x", ">=18", "~18.16", "^1.2.3"
func (v *Version) Satisfies(constraint string) bool {
	constraint = strings.TrimSpace(constraint)

	// Exact match or wildcard
	if !strings.HasPrefix(constraint, ">=") &&
		!strings.HasPrefix(constraint, "~") &&
		!strings.HasPrefix(constraint, "^") {
		c, err := Parse(constraint)
		if err != nil {
			return false
		}

		// Wildcard: 18.x matches any 18.y.z
		if strings.HasSuffix(constraint, ".x") || !strings.Contains(constraint, ".") {
			return v.Major == c.Major
		}

		return v.Compare(c) == 0
	}

	// >= constraint
	if strings.HasPrefix(constraint, ">=") {
		c, err := Parse(strings.TrimPrefix(constraint, ">="))
		if err != nil {
			return false
		}
		return v.Compare(c) >= 0
	}

	// ~ constraint (tilde range: ~1.2.3 := >=1.2.3 <1.3.0)
	if strings.HasPrefix(constraint, "~") {
		c, err := Parse(strings.TrimPrefix(constraint, "~"))
		if err != nil {
			return false
		}

		// Must have same major and minor
		if v.Major != c.Major || v.Minor != c.Minor {
			return false
		}

		// Must be >= patch version
		return v.Patch >= c.Patch
	}

	// ^ constraint (caret range: ^1.2.3 := >=1.2.3 <2.0.0)
	if strings.HasPrefix(constraint, "^") {
		c, err := Parse(strings.TrimPrefix(constraint, "^"))
		if err != nil {
			return false
		}

		// Must have same major version
		if v.Major != c.Major {
			return false
		}

		// Must be >= minor.patch
		if v.Minor < c.Minor {
			return false
		}
		if v.Minor == c.Minor && v.Patch < c.Patch {
			return false
		}

		return true
	}

	return false
}
