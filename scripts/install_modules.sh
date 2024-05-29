#!/bin/bash
set -e

source ./parse_packages.sh

download_modules (){
  local module_path=$1
  echo "Downloading modules : $module_path"
  cd "$module_path"
  go mod download
  cd - > /dev/null # Back to previous directory
}

# parse package list
parse_package_list "$1"

# Iterate over each module path and download
# shellcheck disable=SC2154
for module in "${list[@]}"; do
  download_modules $module
done

echo "Modules downloaded successfully"