#!/bin/bash
set -e

# Install Homebrew if not installed
if ! command -v brew >/dev/null 2>&1; then
    echo "Homebrew not found. Install it from https://brew.sh/ and rerun this script."
    exit 1
fi

# Install dependencies
brew install go tree cloc

# Install skim
go install github.com/cmsolson75/skim/cmd/skim@latest

# Add Go bin to PATH for current session
export PATH="$HOME/go/bin:$PATH"

# Persist PATH update
if ! grep -q 'export PATH="$HOME/go/bin:$PATH"' ~/.zprofile && [ -f ~/.zprofile ]; then
    echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zprofile
elif [ ! -f ~/.zprofile ]; then
    echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zprofile
fi

echo "Setup complete. Run 'source ~/.zprofile' or restart the terminal."
