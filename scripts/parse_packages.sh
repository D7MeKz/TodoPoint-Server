#!/bin/bash

PACKAGES_FILE="workspace.packages.json"
# Function to display usage
usage() {
  echo "Usage: $0 <key>"
  echo "Example: $0 auth"
  exit 1
}

# Function to check if the correct number of arguments is provided
check_arguments() {
  if [ "$#" -ne 1 ]; then
    usage
  fi
}

# Function to check if the JSON file exists
check_file_exists() {
  if [ ! -f "$1" ]; then
    echo "File $1 not found!"
    exit 1
  fi
}

# Function to extract the list for the given key using jq
extract_list() {
  local json_file=$1
  local key=$2
  jq -r --arg key "$key" '.[$key][]' "$json_file"
}

# Main script execution
parse_package_list() {
  check_arguments "$@"

  local json_file=$PACKAGES_FILE
  local key=$1

  check_file_exists "$json_file"

  local list
  list=$(extract_list "$json_file" "$key")

  if [ -z "$list" ]; then
    echo "Key $key not found in the JSON file or the list is empty."
    exit 1
  fi

  echo "$list"
}



