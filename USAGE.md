# unipm Usage Guide

## Quick Start

### 1. Initialize a new project

```bash
unipm init
```

This creates a `devpack.yaml` file:

```yaml
apps:
  - vscode
  - git
  - node
```

### 2. Check system readiness

```bash
unipm doctor
```

Output:
```
OS: darwin
Architecture: amd64

✓ brew: available

All required tools are available.
```

### 3. Preview the installation plan

```bash
unipm plan
```

Output:
```
Plan for darwin:

  vscode → brew install --cask visual-studio-code
  git → brew install git
  node → brew install node

To apply this plan, run 'unipm apply'.
```

### 4. Apply the plan

```bash
unipm apply
```

Or skip confirmation:
```bash
unipm apply --yes
```

Dry run (no actual execution):
```bash
unipm apply --dry-run
```

---

## Configuration

### devpack.yaml

```yaml
apps:
  - vscode      # Visual Studio Code
  - git         # Git
  - node        # Node.js
  - python      # Python
  - docker      # Docker
  - kubectl     # kubectl
  - terraform   # Terraform
```

---

## Available Packages

See the [unipm-registry](https://github.com/Litchi-group/unipm-registry) for the full list of supported packages.

**Current packages (20):**

- Development Tools: vscode, git, vim
- Languages: node, python, go, rust
- Package Managers: npm, yarn
- DevOps: docker, kubectl, terraform
- Utilities: curl, wget, jq, fzf, htop, tmux
- Databases: postgresql, redis

---

## Command Reference

### `unipm init`
Creates a `devpack.yaml` file in the current directory.

**Flags:** None

---

### `unipm doctor`
Checks if required package managers are installed.

**Flags:** None

**Example:**
```bash
$ unipm doctor
OS: darwin
Architecture: amd64

✓ brew: available

All required tools are available.
```

---

### `unipm plan`
Generates an installation plan without executing.

**Flags:** None

**Example:**
```bash
$ unipm plan
Plan for darwin:

  vscode → brew install --cask visual-studio-code
  git → brew install git

To apply this plan, run 'unipm apply'.
```

---

### `unipm apply`
Executes the installation plan.

**Flags:**
- `--dry-run` - Show what would be done without executing
- `-y, --yes` - Skip confirmation prompt

**Example:**
```bash
$ unipm apply --yes
Plan for darwin:

  vscode → brew install --cask visual-studio-code
  git (already installed)

Installing vscode...
  ✓ Installed
Installing git...
  ⊙ Already installed

Done! 1 installed, 1 skipped.
```

---

## Environment Variables

### `UNIPM_REGISTRY_PATH`
Use a local registry instead of GitHub (for development/testing).

```bash
export UNIPM_REGISTRY_PATH=/path/to/local/registry
unipm plan
```

---

## Examples

### Web Development Setup

```yaml
apps:
  - vscode
  - git
  - node
  - npm
  - yarn
  - docker
```

### Python Development

```yaml
apps:
  - vscode
  - git
  - python
  - docker
  - postgresql
```

### DevOps Toolkit

```yaml
apps:
  - git
  - docker
  - kubectl
  - terraform
  - jq
  - curl
  - wget
```

### System Utilities

```yaml
apps:
  - vim
  - tmux
  - htop
  - fzf
  - jq
  - curl
  - wget
```

---

## Troubleshooting

### Package not found
```
✗ mypackage: package not found: mypackage
```

**Solution:** The package is not yet in the registry. Check [available packages](https://github.com/Litchi-group/unipm-registry) or [contribute a new package](https://github.com/Litchi-group/unipm-registry#contributing).

### Provider not available
```
✗ brew: not found
```

**Solution:** Install the required package manager. Run `unipm doctor` for installation instructions.

### Permission denied
```
failed to install: permission denied
```

**Solution:** Some package managers (apt, snap) require sudo. unipm will prompt when needed.

---

## FAQ

**Q: Does unipm replace Homebrew/WinGet/apt?**  
A: No. unipm orchestrates existing package managers. You still need them installed.

**Q: Can I use custom package names?**  
A: Not yet. v0.1 uses the public registry only. Custom registries are planned for future versions.

**Q: Does unipm support version pinning?**  
A: Not in v0.1. This is planned for v0.2.

**Q: Can I use unipm for team environments?**  
A: Yes! Commit `devpack.yaml` to your repo. Team members run `unipm apply` to get the same setup.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to contribute packages or features.
