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

## [Unreleased]

### Planned for v0.2
- Version pinning and lockfiles
- Custom/private registries
- IDE and dotfile synchronization
- Package search command
- More package managers (chocolatey, scoop, etc.)
- More packages in registry

---

[0.1.0]: https://github.com/Litchi-group/unipm/releases/tag/v0.1.0
