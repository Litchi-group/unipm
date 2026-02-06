# Contributing to unipm

Thank you for your interest in contributing to unipm! 🎉

---

## How to Contribute

### 1. Add a Package to the Registry

The easiest way to contribute is by adding new packages to [unipm-registry](https://github.com/Litchi-group/unipm-registry).

#### Steps:

1. Fork the [unipm-registry](https://github.com/Litchi-group/unipm-registry) repository
2. Create a new file: `packages/<package-id>.yaml`
3. Follow the format:

```yaml
id: package-id
name: Package Name
homepage: https://example.com/
providers:
  macos:
    - type: brew        # or brew_cask
      name: package-name
  windows:
    - type: winget
      id: Publisher.PackageName
  linux:
    - type: apt         # or snap
      name: package-name
      classic: true     # snap only, optional
```

4. Test locally:
```bash
export UNIPM_REGISTRY_PATH=/path/to/your/fork
unipm plan
```

5. Submit a Pull Request

---

### 2. Report Bugs

Found a bug? Please [open an issue](https://github.com/Litchi-group/unipm/issues/new) with:

- unipm version (`unipm --version`)
- Operating system and version
- Steps to reproduce
- Expected vs actual behavior

---

### 3. Suggest Features

Have an idea? [Open a feature request](https://github.com/Litchi-group/unipm/issues/new) and describe:

- The problem you're trying to solve
- Your proposed solution
- Any alternatives you've considered

---

### 4. Contribute Code

#### Setup Development Environment

```bash
# Clone the repo
git clone https://github.com/Litchi-group/unipm.git
cd unipm

# Install dependencies
go mod download

# Build
go build -o unipm .

# Run tests
go test ./...
```

#### Make Changes

1. Create a new branch:
```bash
git checkout -b feature/my-feature
```

2. Make your changes

3. Test thoroughly:
```bash
go build -o unipm .
./unipm doctor
./unipm plan
```

4. Commit with a clear message:
```bash
git commit -m "feat: add support for XYZ"
```

Use conventional commit format:
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation
- `refactor:` - Code refactoring
- `test:` - Tests

5. Push and open a PR

---

## Code Guidelines

- Follow Go best practices
- Add comments for complex logic
- Keep functions small and focused
- Write tests for new features

---

## Testing

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./internal/provider

# With coverage
go test -cover ./...
```

---

## Project Structure

```
unipm/
├── cmd/              # CLI commands (init, plan, apply, doctor)
├── internal/
│   ├── config/       # devpack.yaml loader
│   ├── detector/     # OS detection
│   ├── provider/     # Package manager providers
│   ├── registry/     # Registry loader & resolver
│   └── planner/      # Installation planner
├── main.go
├── Makefile
└── README.md
```

---

## Release Process

Releases are automated via GitHub Actions when a tag is pushed:

```bash
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```

---

## Questions?

Feel free to open an issue or reach out to the maintainers.

Thank you for contributing! 🙏
