#!/bin/bash

# Initialize Go workspace
function init_workspace {
    # Check if go.work already exists
    if [ -f go.work ]; then
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

    # Loop through each package and add it to the workspace
    for package in $packages; do
        go work use "$package"
    done

    echo "Common modules installed."
}

# Compare hash of each package in workspace.packages.json
function compare {
    local json_file="workspace.packages.json"
    local needs_install=false

    # Read the package names and hashes from the JSON file
    packages=$(jq -r '.packages[]' $json_file)

    # Compare the current hash of each package directory with the stored hash
    i=0
    for package in $packages; do
        # Calculate the current hash of the package directory contents
        echo $package

        calculated_hash=$(find "$package" -type f -exec sha256sum {} + | sort | sha256sum | awk '{ print $1 }')

        echo  $calculated_hash

        stored_hash=$(jq -r ".hashes[$i]" $json_file)

        # If the hash is empty or mismatched, mark for installation
        if [ -z "$stored_hash" ] || [ "$stored_hash" != "$calculated_hash" ]; then
            echo "Hash is empty or mismatch for $package. Installing common modules..."
            needs_install=true
        fi

        # Update the hash in the temporary JSON file
        jq --arg new_hash "$calculated_hash" ".hashes[$i] = \$new_hash" $json_file > temp.json && mv temp.json $json_file

        i=$((i+1))
    done

    # If any hash mismatched, install the common modules
    if [ "$needs_install" = true ]; then
        install_common_modules
        echo "Hashes updated in workspace.packages.json"
    else
        echo "All hashes match. No need to install common modules."
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
