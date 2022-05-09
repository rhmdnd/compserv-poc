#!/bin/bash

display_description() {
    echo "Generate code from OpenAPI server"
}

print_usage() {
    cmd=$(basename $0)
    echo "Usage: $cmd"
    echo "    -g [string]             Generator name (optional). The default is go-server, but go-gin-server is acceptable."
    echo "    -i [string]             OpenAPI specification input file (optional)."
    echo "    -h                      Display this help message."
}

while getopts "g:i:h" opt; do
    case ${opt} in
        g )
            GENERATOR="$OPTARG"
            ;;
        i )
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

# The go-gin-server and go-server use different arguments for the source code
# path. Let's make sure we use the right one depending on the generator.
if [[ $GENERATOR_NAME == "go-server" ]]; then
  CODE_PATH_PROPERTY="--additional-properties=sourceFolder=compserv"
elif [[ $GENERATOR_NAME == "go-gin-server" ]]; then
  CODE_PATH_PROPERTY="--additional-properties=apiPath=compserv"
else
  echo "Generator $GENERATOR isn't supported by this script."
  print_usage
  exit 1
fi

java -jar $GENERATOR_PATH generate \
  -g $GENERATOR_NAME \
  -i $INPUT \
  --package-name compserv \
  $CODE_PATH_PROPERTY \
  --output $OUTPUT_DIR
