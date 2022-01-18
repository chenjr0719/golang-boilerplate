#!/usr/bin/env bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do SOURCE="$(readlink "$SOURCE")"; done
ROOT_DIR="$(cd -P "$(dirname "$SOURCE")/.." && pwd)"

# Setup RMQ
./hack/docker-compose-up.sh --dev

go test -cover -coverprofile=c.out $ROOT_DIR/pkg/...
go tool cover -html=c.out -o coverage.html
