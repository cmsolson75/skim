#!/bin/bash
set -e

# Update package index
sudo apt update

# Install dependencies
sudo apt install -y golang tree cloc

# Install skim
go install github.com/cmsolson75/skim/cmd/skim@latest

# Add Go bin to PATH for current session
export PATH="$HOME/go/bin:$PATH"

# Persist PATH update
if ! grep -q 'export PATH="$HOME/go/bin:$PATH"' ~/.bashrc; then
    echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
fi

echo "Setup complete. Run 'source ~/.bashrc' or restart the terminal."
