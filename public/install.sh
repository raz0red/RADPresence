#!/usr/bin/env sh
# Installs the latest RAD Presence release for macOS/Linux.
#
#   curl -fsSL https://radpresence.com/install.sh | sh
#
set -eu

REPO="raz0red/RADPresence"
INSTALL_DIR="${RADPRESENCE_INSTALL_DIR:-$HOME/.local/bin}"
BIN_NAME="radpresence"

os=$(uname -s)
arch=$(uname -m)

case "$os" in
  Darwin) platform="darwin" ;;
  Linux) platform="linux" ;;
  *)
    echo "error: unsupported OS: $os" >&2
    echo "RAD Presence supports macOS and Linux via this script. Windows users: download radpresence.exe from https://github.com/${REPO}/releases" >&2
    exit 1
    ;;
esac

case "$arch" in
  x86_64|amd64) goarch="amd64" ;;
  arm64|aarch64) goarch="arm64" ;;
  *)
    echo "error: unsupported architecture: $arch" >&2
    exit 1
    ;;
esac

api_url="https://api.github.com/repos/${REPO}/releases/latest"
echo "Fetching latest release info..."
asset_url=$(curl -fsSL "$api_url" | grep '"browser_download_url"' | grep "${platform}-${goarch}" | head -n1 | cut -d '"' -f4)

if [ -z "$asset_url" ]; then
  echo "error: could not find a release asset for ${platform}-${goarch}" >&2
  exit 1
fi

mkdir -p "$INSTALL_DIR"
tmp_file=$(mktemp)
echo "Downloading $asset_url"
curl -fsSL "$asset_url" -o "$tmp_file"
chmod +x "$tmp_file"
mv "$tmp_file" "$INSTALL_DIR/$BIN_NAME"

echo ""
echo "RAD Presence installed to $INSTALL_DIR/$BIN_NAME"

case ":$PATH:" in
  *":$INSTALL_DIR:"*) ;;
  *)
    echo ""
    echo "NOTE: $INSTALL_DIR is not on your PATH."
    echo "Add this to your shell profile (~/.zshrc, ~/.bashrc, etc.):"
    echo "  export PATH=\"$INSTALL_DIR:\$PATH\""
    ;;
esac

echo ""
echo "Next steps:"
echo "  $BIN_NAME set --username YOUR_RA_USERNAME --apikey YOUR_API_KEY"
echo "  $BIN_NAME run          # test in the foreground first"
echo "  $BIN_NAME install && $BIN_NAME start   # install as a background service"
