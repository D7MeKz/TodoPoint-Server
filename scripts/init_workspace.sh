#!/bin/bash

PACKAGE_FILE="workspace.packages.json"

# Initialize Go workspace
function init_workspace {
    # Check if go.mod already exists
    if [ -f go.work ]; then
        rm go.work
        go work init
    else
        go work init
        echo "Go workspace initialized."
    fi
}

# Install common modules specified in workspace.packages.json
function install_common_modules {
    # Read package names from workspace.packages.json
    packages=$(jq -r '.packages[]' workspace.packages.json)

    # Loop through each package and install
    for package in $packages; do
        go work use "$package"
        if [ "$package" == "." ]; then
          continue
        fi
        go mod download
        if  [ $? -ne 0 ]; then
            echo "Error: Failed to download Go modules"
            exit 1
        else
            echo "Go modules downloaded successfully."
        fi


    done

    echo "Common modules installed."
}

# Function to calculate MD5 hash of a string based on OS
function calculate_md5 {
    case "$(uname -s)" in
        Linux*)     MD5_CMD="md5sum";;
        Darwin*)    MD5_CMD="md5 -r";;
        *)          echo "Unsupported OS"; exit 1;;
    esac
    echo -n "$1" | $MD5_CMD | awk '{ print $1 }'
}

function update_md5() {

  version=$(jq -r '.version' $PACKAGE_FILE)
  packages=$(jq -r '.packages[]' $PACKAGE_FILE | tr -d '\n')
  combined="$version$packages"
  echo "Combined version and packages: $combined"

  # Calculate MD5 hash of the combined string
  md5_hash=$(calculate_md5 "$combined")

  current_hash=$(jq -r '.hash' $PACKAGE_FILE)
  if [ -z "$current_hash" ] || [ "$current_hash" != "$md5_hash" ]; then
        jq --arg md5_hash "$md5_hash" '.hash = $md5_hash' $PACKAGE_FILE > temp.json && mv temp.json $PACKAGE_FILE
        echo "Hash field set with MD5 hash in $PACKAGE_FILE"

        install_common_modules
  else
        echo "No existing hash found. Setting new hash value..."
  fi

}

# Main function
function main {
    # Initialize Go workspace
    # Check if go.mod already exists
    if [ -f go.work ]; then
        echo "Go workspace already initialized."
    else
        go work init
        echo "Go workspace initialized."
    fi
    update_md5
}

# Execute main function
main
