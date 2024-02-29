#!/bin/bash

echo "   __  __           __          _                __"
echo "  / / / /___  _____/ /_  ____ _(_)___  ___  ____/ /"
echo " / / / / __ \\/ ___/ __ \\/ __ \`/ / __ \\/ _ \\/ __  /"
echo "/ /_/ / / / / /__/ / / / /_/ / / / / /  __/ /_/ /  "
echo "\____/_/ /_/\\___/_/ /_/\\__,_/_/_/ /_/\\___/\\__,_(_) "
echo ""

# List of required commands
required_commands=("curl")

# Loop through the list and check each command
missing_commands=()
for cmd in "${required_commands[@]}"; do
  if ! command -v $cmd &>/dev/null; then
    # Command is missing, add it to the list of missing commands
    missing_commands+=($cmd)
  fi
done

# Check if there are any missing commands
if [ ${#missing_commands[@]} -ne 0 ]; then
  echo "The following required command(s) could not be found:"
  for cmd in "${missing_commands[@]}"; do
    echo "- $cmd"
  done
  echo "Please install the missing command(s) and try again."
  exit 1
fi

# Detect operating system
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  INF_OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  INF_OS="darwin"
elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys"* || "$OSTYPE" == "win32" ]]; then
  INF_OS="windows"
else
  echo "Unsupported OS"
  exit 1
fi

# Detect CPU architecture
ARCH=$(uname -m)
if [[ "$ARCH" == "x86_64" ]]; then
  INF_CPU="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  INF_CPU="arm64"
else
  echo "Unsupported CPU architecture"
  exit 1
fi

# Export the variables (optional, useful if sourcing the script)
export INF_OS
export INF_CPU

# Print the variables for verification
echo "Operating System: $INF_OS"
echo "CPU Architecture: $INF_CPU"

# GitHub repository details
REPO="KenshiTech/unchained"
GITHUB_API_URL="https://api.github.com/repos/$REPO/releases/latest"

# Fetch the latest release tag
TAG=$(curl -s $GITHUB_API_URL | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$TAG" ]; then
  echo "Failed to fetch the latest release tag"
  exit 1
fi

# Construct the filename based on the OS, ARCH, and for Windows, add .exe
FILENAME="unchained.$INF_OS.$INF_CPU"
if [[ "$INF_OS" == "windows" ]]; then
  FILENAME="$FILENAME.exe"
fi

# Construct the download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$TAG/$FILENAME"

# Define the installation path
if [ -z "$INSTALL_PATH" ]; then
  if [[ "$INF_OS" == "linux" || "$INF_OS" == "darwin" ]]; then
    INSTALL_PATH="/usr/local/bin"
  elif [[ "$INF_OS" == "windows" ]]; then
    # For simplicity, we're using a generic path; adjust as necessary for your environment
    # Attempt to use a Unix-like install path if available and writable
    if [[ -d "/usr/local/bin" && -w "/usr/local/bin" ]]; then
      INSTALL_PATH="/usr/local/bin"
    else
      # Default to the current working directory if /usr/local/bin isn't suitable
      echo "No /usr/local/bin available"
      INSTALL_PATH=$(pwd)
    fi
  fi
fi

# Download the latest release
echo "Downloading the latest release..."
curl -sS -L $DOWNLOAD_URL -o /tmp/$FILENAME

if [ $? -ne 0 ]; then
  echo "Failed to download the file."
  exit 1
fi

# Determine if we can install automatically on Windows environments
if [[ "$INF_OS" == "windows" ]]; then
  # Execute installation based on determined action
  if [[ -n "$WSL_DISTRO_NAME" || -n "$CYGWIN" ]]; then
    chmod +x /tmp/$FILENAME
    mv /tmp/$FILENAME $INSTALL_PATH
    echo "Unchained has been installed to $INSTALL_PATH"
  else
    echo "Downloaded $FILENAME. Please move it to your desired location and add it to your PATH."
  fi
else
  # Non-Windows installation (Linux, macOS) proceeds as before
  chmod +x /tmp/$FILENAME
  sudo mv /tmp/$FILENAME $INSTALL_PATH/unchained
  if [ $? -ne 0 ]; then
    echo "Failed to install Unchained."
  else
    echo "Unchained $TAG has been installed to $INSTALL_PATH"
  fi
fi
