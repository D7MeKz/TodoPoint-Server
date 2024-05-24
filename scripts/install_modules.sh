#!/bin/bash
set -e

MODULES=(
      "/modules/common"
      "/modules/database/d7mysql"
      "/modules/database/d7redis"
)

download_modules (){
  local module_path=$1
  echo "Downloading modules : $module_path"
  cd $module_path
  go mod download
  cd - > /dev/null # Back to previous directory
}

# Iterate over each module path and download
for module in "${MODULES[@]}"; do
  download_modules $module
done

echo "Modules downloaded successfully"