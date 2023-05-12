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

GO_PLUGIN_PATH=$(command -v protoc-gen-gogo 2>&1)
if [[ "$GO_PLUGIN_PATH" == "" ]]; then
  echo "'protoc-gen-go' tool is required but missing." >&2
  exit 1
fi

protodep up

# go
(
  rm -rf golang && mkdir -p golang
  $PROTOC_PATH \
    -I . \
    -I ./protodep \
    --gogo_out=plugins=grpc,paths=source_relative:golang \
    ./chargehive/**/*.proto
)

# php - remove gogo registration from php
(
  rm -rf tmp_php_proto && mkdir -p tmp_php_proto
  cp -R chargehive tmp_php_proto
  gsed -i'' -E 's#^import "gogoproto/gogo.proto";$##g' tmp_php_proto/chargehive/**/*.proto
  gsed -i'' -E 's#^option \(gogoproto\..+##g' tmp_php_proto/chargehive/**/*.proto

  rm -rf php && mkdir -p php
  $PROTOC_PATH \
    -I ./tmp_php_proto \
    -I ./protodep \
    --php_out=php \
    ./tmp_php_proto/chargehive/**/*.proto

  # cleanup
  rm -Rf tmp_php_proto
)
