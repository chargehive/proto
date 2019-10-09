#!/bin/bash

set -e

# dependency check
PROTOC_PATH=$(command -v protoc 2>&1)
if [[ "$PROTOC_PATH" == "" ]]; then
  echo "'protoc' tool is required but missing." >&2
  exit 1
fi

# go
rm -rf golang
mkdir -p golang

$PROTOC_PATH -I . \
  -I $GOPATH/src \
  --gogo_out=plugins=grpc,paths=source_relative:golang \
  chargehive/**/*.proto

# init
mkdir -p genproto
cp -R chargehive genproto

# php
rm -rf php
mkdir -p php/src
sed -i '' -E 's#^import "github.com/gogo/protobuf/gogoproto/gogo.proto";$##g' genproto/chargehive/**/*.proto
sed -i '' -E 's#^option \(gogoproto\..+##g' genproto/chargehive/**/*.proto

$PROTOC_PATH -I genproto \
  -I $GOPATH/src \
  --php_out=php \
  genproto/chargehive/**/*.proto

# cleanup
rm -Rf genproto
