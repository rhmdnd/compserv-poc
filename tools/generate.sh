#!/bin/bash

display_description() {
    echo "Generate code from OpenAPI server"
}

print_usage() {
    cmd=$(basename $0)
    echo "Usage: $cmd"
    echo "    -g [string]             Generator name (optional)."
    echo "    -i [string]             OpenAPI specification input file (optional)."
    echo "    -h                      Display this help message."
}

while getopts "g:i:h" opt; do
    case ${opt} in
        g )
            GENERATOR="$OPTARG"
            ;;
        g )
            SPECIFICATION="$OPTARG"
            ;;
        h )
            display_description
            print_usage
            exit 0
            ;;
        \? )
            print_usage
            exit 1
            ;;
    esac
done
OUTPUT_DIR=/vagrant/
GENERATOR_NAME=${GENERATOR:-go-server}
GENERATOR_PATH=$HOME/openapi-generator-cli-6.0.0-beta.jar
INPUT=${SPECIFICATION:-https://raw.githubusercontent.com/EasyDynamics/oscal-rest/develop/openapi.yaml}

java -jar $GENERATOR_PATH generate \
  -g $GENERATOR_NAME \
  -i $INPUT \
  --package-name compserv \
  --additional-properties=sourceFolder=compserv \
  --output $OUTPUT_DIR
