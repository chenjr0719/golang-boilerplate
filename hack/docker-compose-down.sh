#!/usr/bin/env bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do SOURCE="$(readlink "$SOURCE")"; done
ROOT_DIR="$(cd -P "$(dirname "$SOURCE")/.." && pwd)"
DEPLOYMENT_DIR=$ROOT_DIR/deployments

POSITIONAL=()
while [[ $# -gt 0 ]]; do
    key="$1"

    case $key in
    --dev)
        DEV=TRUE
        shift # past argument
        ;;
    --purge)
        PURGE=TRUE
        shift # past argument
        ;;
    *)                     # unknown option
        POSITIONAL+=("$1") # save it in an array for later
        shift              # past argument
        ;;
    esac
done

set -- "${POSITIONAL[@]}" # restore positional parameters

if [[ ! -z "$DEV" ]] && [[ "$DEV" == "TRUE" ]]; then
    COMPOSE_ARGS="-f $DEPLOYMENT_DIR/docker-compose.dev.yml"
else
    COMPOSE_ARGS=""
fi

if [[ ! -z "$PURGE" ]] && [[ "$PURGE" == "TRUE" ]]; then
    DOWN_ARGS="--remove-orphans --volumes"
else
    DOWN_ARGS=""
fi

docker-compose \
    --project-directory $ROOT_DIR \
    -f $DEPLOYMENT_DIR/docker-compose.yml \
    $COMPOSE_ARGS \
    down $DOWN_ARGS
