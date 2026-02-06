# unipm

> Universal package manager orchestrator for Windows, macOS, and Linux.

**unipm** is a cross-platform CLI that lets you define your development environment once
and apply it consistently across different operating systems by orchestrating native package managers.

---

## ✨ Why unipm?

Developers using multiple operating systems often face the same problems:

- Different package managers per OS (brew, winget, apt, etc.)
- Different package names for the same tool
- Repetitive and error-prone environment setup
- No simple way to declare a portable dev environment

**unipm solves this by acting as an orchestrator**, not a replacement.

You write one config file.
unipm figures out how to install it on your OS.

---

## 🧠 Core Concept

- unipm does **not** replace Homebrew, WinGet, or apt
- unipm **invokes** them in a unified, declarative way
- One logical package name → OS-specific install command

```
vscode
 ├─ macOS   → brew install --cask visual-studio-code
 ├─ Windows → winget install Microsoft.VisualStudioCode
 └─ Linux   → snap install code --classic
```

---

## 🚀 Features (v0.1)

- Cross-platform CLI (Windows / macOS / Linux)
- Declarative YAML-based configuration
- Installation plan preview (dry-run by default)
- Optional execution of the plan
- No server, no daemon, no runtime dependencies
- Single static binary

---

## 📦 Installation

### From Releases (Recommended)

Download the latest release for your platform:

**macOS (Intel)**
```bash
curl -L https://github.com/Litchi-group/unipm/releases/latest/download/unipm-darwin-amd64 \
  -o /usr/local/bin/unipm
chmod +x /usr/local/bin/unipm
```

**macOS (Apple Silicon)**
```bash
curl -L https://github.com/Litchi-group/unipm/releases/latest/download/unipm-darwin-arm64 \
  -o /usr/local/bin/unipm
chmod +x /usr/local/bin/unipm
```

**Linux (amd64)**
```bash
curl -L https://github.com/Litchi-group/unipm/releases/latest/download/unipm-linux-amd64 \
  -o /usr/local/bin/unipm
chmod +x /usr/local/bin/unipm
```

**Windows (amd64)**
Download from [Releases](https://github.com/Litchi-group/unipm/releases) and add to PATH.

### From Source
```bash
git clone https://github.com/Litchi-group/unipm.git
cd unipm
go install .
```

---

## 🛠 Usage

### 1. Initialize a config
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

---

### 2. Preview the installation plan
```bash
unipm plan
```

---

### 3. Apply the plan
```bash
unipm apply
```

Options:
- `--dry-run`
- `--yes`

---

### 4. Check system readiness
```bash
unipm doctor
```

---

## 📜 License

MIT License
