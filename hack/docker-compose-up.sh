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
    --dev-apps)
        DEV_APPS=TRUE
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

if [[ ! -z "$DEV_APPS" ]] && [[ "$DEV_APPS" == "TRUE" ]]; then
    COMPOSE_PROFILE="--profile dev-apps"
else
    COMPOSE_PROFILE=""
fi

docker-compose \
    --project-directory $ROOT_DIR \
    -f $DEPLOYMENT_DIR/docker-compose.yml \
    $COMPOSE_ARGS \
    $COMPOSE_PROFILE \
    up -d
