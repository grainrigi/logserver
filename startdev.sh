#!/bin/bash

set -e
dir=$(cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)
cd "$dir"

export COMPOSE_PROJECT_NAME=logserver

# docker-composeの存在確認
if ! which docker-compose > /dev/null 2>&1; then
  if ! docker compose version > /dev/null 2>&1; then
    echo 'Neither "docker-compose" nor "docker compose" is available. Please install one of them.'
    exit 1
  else
    DOCKER_COMPOSE='docker compose'
  fi
else
  DOCKER_COMPOSE=docker-compose
fi

if [ "$1" == "down" ]; then
  cd .devcontainer
  ${DOCKER_COMPOSE} down
  exit
fi


./.devcontainer/initenv
cd .devcontainer

${DOCKER_COMPOSE} build
${DOCKER_COMPOSE} up -d
${DOCKER_COMPOSE} exec dev /bin/bash -i
