#!/bin/bash

set -e

# dependency check
PROTOC_PATH=$(command -v protoc 2>&1)
if [[ "$PROTOC_PATH" == "" ]]; then
  echo "'protoc' tool is required but missing." >&2
  exit 1
fi

# go
rm -rf golang && mkdir -p golang

$PROTOC_PATH \
  -I "$GOPATH"/pkg/mod \
  -I . \
  --gogo_out=plugins=grpc,paths=source_relative:golang \
  chargehive/**/*.proto

# php - remove gogo registration from php
mkdir -p genproto
cp -R chargehive genproto
rm -rf php && mkdir -p php
gsed -i'' -E 's#^import "github.com/gogo/protobuf([^/]+)?/gogoproto/gogo.proto";$##g' genproto/chargehive/**/*.proto
gsed -i'' -E 's#^option \(gogoproto\..+##g' genproto/chargehive/**/*.proto

$PROTOC_PATH \
  -I "$GOPATH"/pkg/mod \
  -I ./genproto \
  --php_out=php \
  genproto/chargehive/**/*.proto

# cleanup
rm -Rf genproto
