#!/usr/bin/env bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do SOURCE="$(readlink "$SOURCE")"; done
ROOT_DIR="$(cd -P "$(dirname "$SOURCE")/.." && pwd)"
DEPLOYMENT_DIR=$ROOT_DIR/deployments

docker-compose \
    --project-directory $ROOT_DIR \
    -f $DEPLOYMENT_DIR/docker-compose.yml \
    build
