#!/bin/bash

if [ -z "$LOCAL_WORKSPACE_FOLDER" ]; then
  echo "This script must be executed in a Dev Container."
  exit 1
fi

set -e
dir=$(cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)
cd "$dir"

export COMPOSE_PROJECT_NAME=logserver-test

(cd .. && go build .)
echo 'Preparing test environment...'
chronic sudo -E docker compose up -d --build
sudo -E docker compose exec tester sh -c 'chronic yarn && yarn test'
chronic sudo -E docker compose down
