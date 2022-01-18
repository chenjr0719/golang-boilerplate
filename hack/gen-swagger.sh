#!/usr/bin/env bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do SOURCE="$(readlink "$SOURCE")"; done
ROOT_DIR="$(cd -P "$(dirname "$SOURCE")/.." && pwd)"

go install github.com/swaggo/swag/cmd/swag@latest
swag init \
    --dir $ROOT_DIR/cmd/apiserver \
    --output $ROOT_DIR/docs \
    --parseDependency \
    --parseInternal
