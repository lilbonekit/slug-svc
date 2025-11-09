#!/usr/bin/env bash

set -e

PROJECT_DIR=$(pwd)
GENERATED="${PROJECT_DIR}/resources"
OPENAPI_DIR="${PROJECT_DIR}/docs/web_deploy"
PACKAGE_NAME=resources

function printHelp {
    echo "usage: ./generate.sh [<flags>]
    Script to generate Go SDK from OpenAPI spec.

    Flags:
      --package PACKAGE             package name of generated stuff (default: 'resources')
      -p, --path-to-generate PATH   output dir (default: resources)
      -i, --input OPENAPI_DIR       path to dir with openapi.yaml (default: docs/web_deploy)
      -h, --help                    show this help message
    "
}

function parseArgs {
    while [[ -n "$1" ]]; do
        case "$1" in
            -h | --help)
                printHelp && exit 0
                ;;
            -p | --path-to-generate)
                shift
                [[ ! -d $1 ]] && echo "path $1 does not exist or not a dir" && exit 1
                GENERATED=$1
                ;;
            --package)
                shift
                [[ -z "$1" ]] && echo "package name not specified" && exit 1
                PACKAGE_NAME=$1
                ;;
            -i | --input)
                shift
                [[ ! -f "$1/openapi.yaml" ]] && echo "file openapi.yaml does not exist in $1 or not a file" && exit 1
                OPENAPI_DIR=$1
                ;;
        esac
        shift
    done
}

function generate {
    echo "📦 Building docs..."
    (cd docs && npm run build)

    echo "🧹 Cleaning old generated files..."
    rm -rf "${GENERATED:?}/"*

    echo "🚀 Generating Go SDK from OpenAPI spec..."
    docker run --rm \
      -v "${OPENAPI_DIR}":/local/openapi \
      -v "${GENERATED}":/local/generated \
      openapitools/openapi-generator-cli generate \
      -i /local/openapi/openapi.yaml \
      -g go \
      -o /local/generated \
      --skip-validate-spec \
      --additional-properties=packageName=${PACKAGE_NAME},isGoSubmodule=false,enumClassPrefix=true,typeMappings=type=ResourceType,date-time=time.Time

    # 🧹 Remove redundant go.mod/go.sum just in case
    if [[ -f "${GENERATED}/go.mod" ]]; then
        echo "⚠️  Removing redundant go.mod from resources/"
        rm -f "${GENERATED}/go.mod"
    fi
    if [[ -f "${GENERATED}/go.sum" ]]; then
        rm -f "${GENERATED}/go.sum"
    fi

    echo "🧹 Formatting generated code..."
    if command -v goimports >/dev/null 2>&1; then
        goimports -w ${GENERATED}
    else
        go fmt ${GENERATED}/...
    fi

    echo "✅ Generation complete. SDK is in: ${GENERATED}"
}

parseArgs "$@"
generate
