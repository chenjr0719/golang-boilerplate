#!/usr/bin/env bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do SOURCE="$(readlink "$SOURCE")"; done
ROOT_DIR="$(cd -P "$(dirname "$SOURCE")/.." && pwd)"
BUILD_DIR=$ROOT_DIR/build

docker build $ROOT_DIR -f $BUILD_DIR/worker/Dockerfile -t golang-boilerplate-worker:latest
