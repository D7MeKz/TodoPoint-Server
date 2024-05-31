#!/bin/bash

# Initialize Go workspace
function init_workspace {
    # Check if go.mod already exists
    if [ -f go.mod ]; then
        echo "Go workspace already initialized."
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
    done

    echo "Common modules installed."
}

# Compare hash of workspace.packages.json
function compare {
    local current_hash=$(jq -r '.hash' workspace.packages.json)
    local calculated_hash=$(jq -c '.packages' workspace.packages.json | sha256sum | awk '{ print $1 }')

    # Check if hash matches
    if [ "$current_hash" != "$calculated_hash" ]; then
        echo "Hash mismatch, installing common modules..."
        install_common_modules

        # Update the hash in workspace.packages.json
        jq --arg new_hash "$calculated_hash" '.hash = $new_hash' workspace.packages.json > temp.json && mv temp.json workspace.packages.json
        echo "Hash updated in workspace.packages.json"
    else
        echo "Hashes match, no need to install common modules."
    fi
}

# Main function
function main {
    # Initialize Go workspace
    init_workspace

    # Compare and possibly install common modules
    compare
}

# Execute main function
main
