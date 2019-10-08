#!/bin/bash

set -e

# dependency check
PROTOC_PATH=$(command -v protoc 2>&1)
if [[ "$PROTOC_PATH" == "" ]]; then
  echo "'protoc' tool is required but missing." >&2
  exit 1
fi

# init
mkdir genproto
cp proto/*.proto genproto/

# go
rm -rf golang
mkdir -p golang/chtype
sed -i '' "s#github.com/chargehive/proto/proto/##g" genproto/*.proto
$PROTOC_PATH -I genproto \
  -I $GOPATH/src \
  --gogo_out=plugins=grpc,paths=source_relative:golang/chtype \
  genproto/*.proto

# php
rm -rf php
mkdir -p php/src
sed -i '' -E 's#^import "github.com/gogo/protobuf/gogoproto/gogo.proto";$##g' genproto/*.proto
sed -i '' -E 's#^option \(gogoproto\..+##g' genproto/*.proto
$PROTOC_PATH -I genproto \
  -I $GOPATH/src \
  --php_out=php/src \
  genproto/*.proto

# cleanup
rm -Rf genproto
