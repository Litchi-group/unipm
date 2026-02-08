# Changelog

All notable changes to unipm will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [0.1.0] - 2026-02-06

### Initial Release 🎉

First public release of unipm - Universal package manager orchestrator.

### Added

#### CLI Commands
- `unipm init` - Initialize devpack.yaml configuration
- `unipm doctor` - Check system requirements and package manager availability
- `unipm plan` - Generate installation plan preview
- `unipm apply` - Execute installation plan
  - `--dry-run` flag - Preview without executing
  - `--yes` flag - Skip confirmation prompt

#### Core Features
- Cross-platform support: Windows, macOS, Linux
- OS detection with Linux distribution identification
- Provider system for package managers:
  - Homebrew (macOS) - `brew`, `brew_cask`
  - WinGet (Windows) - `winget`
  - APT (Debian/Ubuntu) - `apt`
  - Snap (Linux) - `snap`
- Registry system:
  - GitHub-based package registry
  - Local cache (24-hour TTL)
  - Local fallback support (`UNIPM_REGISTRY_PATH`)
- Installation planner:
  - Detect already-installed packages
  - Skip installed packages automatically
  - Real-time progress output

#### Packages (20)
- **Development Tools**: vscode, git, vim
- **Languages & Runtimes**: node, python, go, rust
- **Package Managers**: npm, yarn
- **DevOps & Cloud**: docker, kubectl, terraform
- **Utilities**: curl, wget, jq, fzf, htop, tmux
- **Databases**: postgresql, redis

#### Build & Distribution
- Cross-platform builds:
  - macOS (Intel + Apple Silicon)
  - Windows (amd64)
  - Linux (amd64 + arm64)
- Makefile with build targets
- GitHub Actions automated releases
- Single static binaries (no dependencies)

#### Documentation
- README with quick start guide
- USAGE.md with detailed examples
- CONTRIBUTING.md for contributors
- Package registry documentation

### Technical Details
- Language: Go 1.22+
- CLI Framework: cobra v1.10.2
- Config Format: YAML v3
- No runtime dependencies
- Binary sizes: 8.7-9.4 MB

---

## [0.1.2] - 2026-02-09

### Added

#### Version Management
- Version parsing and comparison (`internal/version`)
- Support for semantic versioning constraints:
  - Exact: `node@18`, `python@3.10`
  - Wildcard: `node@18.x`
  - Range: `~18.16`, `^1.2.3`, `>=2.0`
- `PackageSpec` with optional version field

#### Profile System
- Profile support in `devpack.yaml`:
  ```yaml
  profiles:
    web:
      - node
      - vscode
    python:
      - python
      - pycharm
  ```
- `--profile` flag for `plan` and `apply` commands
- Profile-specific package lists

#### Export/Import
- `unipm export [file]` - Export installed packages to devpack.yaml
- `unipm import [file]` - Import and install from devpack.yaml
- `ListInstalled()` interface for all providers
- Cross-provider package detection

#### Configuration System
- Global config: `~/.unipm/config.yaml`
- Registry URL customization
- Cache TTL configuration
- Log level settings

#### Package Management Commands
- `unipm list` - List installed packages
- `unipm search <query>` - Search registry
- `unipm info <package>` - Show package details
- `unipm update` - Update package managers
- `unipm remove <package>` - Remove packages

#### Dependency Resolution
- Topological sort for dependency ordering
- Circular dependency detection with cycle path
- Diamond dependency handling
- `GetDependencyTree()` for visualization

#### Error Handling
- Structured error types:
  - `NotFoundError` - Package not found
  - `CircularDependencyError` - Dependency cycle
  - `ProviderUnavailableError` - Missing provider
  - `NetworkError` - Registry access issues
  - `ConfigError` - Configuration problems
- User-friendly error messages with suggestions

#### Infrastructure
- Test infrastructure:
  - Test helpers (`internal/testing/helpers.go`)
  - MockRegistry for testing
  - Detector tests (8 tests)
- Logging system:
  - `internal/logger` with DEBUG/INFO/WARN/ERROR levels
  - `--verbose` flag for debug output
  - Structured logging in providers
- CI/CD:
  - `.github/workflows/test.yml` - Automated testing
  - `.github/workflows/lint.yml` - Code linting
  - golangci-lint configuration
- Code refactoring:
  - `BaseProvider` with common functionality
  - Registry interface (`RegistryInterface`)
  - Reduced code duplication by ~15%

#### Security
- SHA256 checksum support for packages
- Checksum verification (`VerifyChecksum`)
- `checksum` field in package registry

### Changed
- Improved provider implementations:
  - Extracted common patterns to `BaseProvider`
  - Added logging to all provider operations
  - Simplified BrewProvider and WinGetProvider
- Enhanced config schema:
  - Added `profiles` section
  - Support for package version specs (`name@version`)
- Version bumped to 0.1.2 in `cmd/root.go`

### Fixed
- Dependency resolution edge cases
- Error propagation in command handlers

### Internal
- Registry interface for better testability
- Improved error propagation
- Code organization and documentation
- 70%+ test coverage target (in progress)

---

## [0.1.1] - 2026-02-07

### Added
- Dependency resolution with topological sort
- Management commands: list, search, info, update, remove
- index.yaml support for package search
- Provider Remove() method

### Removed
- Custom registry support (UNIPM_REGISTRY_PATH, file:// URLs)

---

## [Unreleased]

### Planned for v0.2
- Test coverage 80%+
- Concurrency improvements
- Progress bar for installations
- godoc documentation
- More packages in registry (20 → 50+)

---

[0.1.2]: https://github.com/Litchi-group/unipm/releases/tag/v0.1.2
[0.1.1]: https://github.com/Litchi-group/unipm/releases/tag/v0.1.1
[0.1.0]: https://github.com/Litchi-group/unipm/releases/tag/v0.1.0
