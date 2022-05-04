#!/bin/bash

display_description() {
    echo "Generate API server code for an OpenAPI specification."
    echo ""
    echo "NOTE: This script assumes that you have the 'podman' binary installed"
    echo ""
}

print_usage() {
    cmd=$(basename $0)
    echo "Usage:"
    echo "    $cmd -h                      Display this help message."
    echo "    $cmd -i [string]             The file path to an OpenAPI specification file as input."
    exit 0
}

while getopts ":i:h" opt; do
    case ${opt} in
        i )
            API_SPEC="$OPTARG"
            ;;
        h )
            display_description
            print_usage
            ;;
        \? )
            print_usage
            ;;
    esac
done

SCRIPT_DIR=$(dirname ${BASH_SOURCE})
UPSTREAM_SPEC="https://raw.githubusercontent.com/EasyDynamics/oscal-rest/develop/openapi.yaml"

if [[ -z $API_SPEC ]]; then
  echo "No OpenAPI specification was provided using -i."
  echo "Using $UPSTREAM_SPEC"
  curl -o $SCRIPT_DIR/openapi.yaml $UPSTREAM_SPEC
  API_SPEC=$SCRIPT_DIR/openapi.yaml
else
  cp $API_SPEC $SCRIPT_DIR/openapi.yaml
fi

# This is a workaround so that the openapi-generator-cli tool can write the
# generated go code to /local, which is mounted from the host file system.
podman unshare chown 1000:1000 ./

podman run --rm \
  --volume "${PWD}:/local:Z" \
  openapitools/openapi-generator-cli generate \
  --additional-properties=generateInterfaces=true \
  --additional-properties=packageVersion=0.1.0 \
  --input-spec /local/openapi.yaml \
  --generator-name go-gin-server \
  --output /local/
