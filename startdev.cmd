@echo off

cd %~dp0

set COMPOSE_PROJECT_NAME=logserver

if "%1"=="stop" (
  cd .devcontainer
  docker-compose down
  exit
)

call .devcontainer\initenv.cmd
cd .devcontainer

docker-compose build
docker-compose up -d
docker-compose exec dev /bin/bash -i