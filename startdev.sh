#!/bin/bash

set -e
dir=$(cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)
cd "$dir"

export COMPOSE_PROJECT_NAME=logserver

if [ "$1" = stop ]; then
  cd .devcontainer
  docker-compose down
  exit
fi

./.devcontainer/initenv
cd .devcontainer

docker-compose build
docker-compose up -d
docker-compose exec dev /bin/bash -i
