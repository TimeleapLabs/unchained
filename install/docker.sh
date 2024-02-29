#!/bin/bash

echo "   __  __           __          _                __"
echo "  / / / /___  _____/ /_  ____ _(_)___  ___  ____/ /"
echo " / / / / __ \\/ ___/ __ \\/ __ \`/ / __ \\/ _ \\/ __  /"
echo "/ /_/ / / / / /__/ / / / /_/ / / / / /  __/ /_/ /  "
echo "\____/_/ /_/\\___/_/ /_/\\__,_/_/_/ /_/\\___/\\__,_(_) "
echo ""

# List of required commands
required_commands=("unzip" "curl" "diff")

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

# GitHub repository details
REPO="KenshiTech/unchained"
GITHUB_API_URL="https://api.github.com/repos/$REPO/releases/latest"

# Fetch the latest release tag
TAG=$(curl -s $GITHUB_API_URL | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$TAG" ]; then
  echo "Failed to fetch the latest release tag"
  exit 1
fi

# Construct the filename based on the latest TAG
FILENAME="unchained-$TAG-docker.zip"

# Construct the download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$TAG/$FILENAME"

# Download the latest release
echo "Downloading the latest release..."
curl -sS -L $DOWNLOAD_URL -o /tmp/$FILENAME

# Check if the download was successful
if [ $? -eq 0 ]; then
  echo "Download completed. Unpacking..."
else
  echo "Failed to download the file."
  exit 1
fi

# Unpack the zip file to the current directory
unzip -o -q /tmp/$FILENAME -d /tmp
echo "Unpacking completed."

# Check if the conf.worker.yaml.template has changed
if [ -f "conf.worker.yaml.template" ]; then
  DIFF=$(diff conf.worker.yaml.template /tmp/unchained-$TAG-docker/conf.worker.yaml.template)
  if [ ! -z "$DIFF" ]; then
    echo "WARNING: The content of the config template has changed in the new release. Please review and update your configuration accordingly."
  fi
fi

if [ -z "$INSTALL_PATH" ]; then
  INSTALL_PATH=$(pwd)
fi

# Move all files from the temporary directory to the current directory
mv /tmp/unchained-$TAG-docker/* $INSTALL_PATH

# Clean up the temporary extraction directory
rm -rf /tmp/unchained-$TAG-docker

echo "Unchained $TAG has been installed to $INSTALL_PATH"
echo ""
echo "┌───────────────────────────────────────────────────────────────┐"
echo "│ To finish the installation, make the required changes to your │"
echo "│ config file in conf/conf.NODE_TYPE.yaml and run:              │"
echo "│                                                               │"
echo "│ ./unchained.sh [node] pull                                    │"
echo "│ ./unchained.sh [node] up -d                                   │"
echo "└───────────────────────────────────────────────────────────────┘"
