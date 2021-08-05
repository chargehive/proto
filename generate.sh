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
rm -rf tmp_php_proto && mkdir -p tmp_php_proto
cp -R chargehive tmp_php_proto
rm -rf php && mkdir -p php
gsed -i'' -E 's#^import "github.com/gogo/protobuf([^/]+)?/gogoproto/gogo.proto";$##g' tmp_php_proto/chargehive/**/*.proto
gsed -i'' -E 's#^option \(gogoproto\..+##g' tmp_php_proto/chargehive/**/*.proto

$PROTOC_PATH \
  -I "$GOPATH"/pkg/mod \
  -I ./tmp_php_proto \
  --php_out=php \
  tmp_php_proto/chargehive/**/*.proto

# cleanup
rm -Rf tmp_php_proto
