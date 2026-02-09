#!/bin/bash
set -e

# unipm installer - universal cross-platform installer
# Usage: curl -fsSL https://raw.githubusercontent.com/Litchi-group/unipm/main/install.sh | bash

REPO="Litchi-group/unipm"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="unipm"

echo "🚀 unipm installer"
echo ""

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
    darwin*)
        OS="darwin"
        ;;
    linux*)
        OS="linux"
        ;;
    mingw* | msys* | cygwin*)
        OS="windows"
        ;;
    *)
        echo "❌ Unsupported OS: $OS"
        exit 1
        ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64 | amd64)
        ARCH="amd64"
        ;;
    arm64 | aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

PLATFORM="${OS}-${ARCH}"
echo "📦 Detected platform: $PLATFORM"

# Determine binary name
if [ "$OS" = "windows" ]; then
    BINARY_NAME="unipm.exe"
fi

DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/unipm-${PLATFORM}"
if [ "$OS" = "windows" ]; then
    DOWNLOAD_URL="${DOWNLOAD_URL}.exe"
fi

echo "📥 Downloading from: $DOWNLOAD_URL"
echo ""

# Download to temp
TEMP_FILE=$(mktemp)
if ! curl -fsSL "$DOWNLOAD_URL" -o "$TEMP_FILE"; then
    echo "❌ Download failed"
    rm -f "$TEMP_FILE"
    exit 1
fi

# Make executable
chmod +x "$TEMP_FILE"

# Install
if [ "$OS" = "windows" ]; then
    # Windows: install to user directory
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
    mv "$TEMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
    echo "✅ Installed to: $INSTALL_DIR/$BINARY_NAME"
    echo ""
    echo "⚠️  Add to PATH: export PATH=\"\$HOME/bin:\$PATH\""
else
    # macOS/Linux: use sudo
    echo "🔐 Installing to $INSTALL_DIR (requires sudo)"
    if sudo mv "$TEMP_FILE" "$INSTALL_DIR/$BINARY_NAME"; then
        echo "✅ Installed to: $INSTALL_DIR/$BINARY_NAME"
    else
        echo "❌ Installation failed (permission denied)"
        echo ""
        echo "Alternative: Install to user directory"
        INSTALL_DIR="$HOME/.local/bin"
        mkdir -p "$INSTALL_DIR"
        mv "$TEMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
        echo "✅ Installed to: $INSTALL_DIR/$BINARY_NAME"
        echo ""
        echo "⚠️  Add to PATH:"
        echo "   echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.bashrc"
        echo "   source ~/.bashrc"
    fi
fi

echo ""
echo "🎉 unipm installed successfully!"
echo ""

# Verify installation
if command -v unipm >/dev/null 2>&1; then
    echo "📍 Location: $(command -v unipm)"
    echo "📌 Version: $(unipm --version 2>/dev/null || echo 'unknown')"
    echo ""
    echo "Run 'unipm --help' to get started!"
else
    echo "⚠️  unipm is installed but not in PATH"
    echo "   Add $INSTALL_DIR to your PATH or restart your terminal"
fi
