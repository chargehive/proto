#!/bin/bash

set -e

# dependency check
PROTOC_PATH=$(command -v protoc 2>&1)
if [[ "$PROTOC_PATH" == "" ]]; then
  echo "'protoc' tool is required but missing." >&2
  exit 1
fi

VER=$(protoc --version)
EXPECT="libprotoc 3.21.12"
if [[ "$VER" != "$EXPECT" ]]; then
  echo "'protoc' tool is version [$VER], expected [$EXPECT]." >&2
  #exit 1 # this is not a critical error
fi

GO_PLUGIN_PATH=$(command -v protoc-gen-go-enums 2>&1)
if [[ "$GO_PLUGIN_PATH" == "" ]]; then
  echo "'protoc-gen-go' tool is required but missing." >&2
  exit 1
fi

protodep up

rm -rf golang && mkdir -p golang
rm -rf php && mkdir -p php
$PROTOC_PATH \
  -I . \
  -I ./protodep \
  --php_out=php \
  --go-enums_out=paths=source_relative,include_nested=true:golang \
  --go_out=paths=source_relative:golang \
  ./chargehive/**/*.proto
