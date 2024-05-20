#!/bin/bash
set -e

MODULES=(
      "/modules/common"
      "/modules/database/d7mysql"
      "/modules/database/d7redis"
)

if [ -f go.work ]; then
  echo "go.work file found, skipping module download"
  rm go.work
fi

go work init

for module in "${MODULES[@]}"; do
  echo "Downloading modules : $module"
  go work use $module
done
