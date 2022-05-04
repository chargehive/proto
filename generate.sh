#!/bin/bash

set -e

# dependency check
PROTOC_PATH=$(command -v protoc 2>&1)
if [[ "$PROTOC_PATH" == "" ]]; then
  echo "'protoc' tool is required but missing." >&2
  exit 1
fi

VER=$(protoc --version)
EXPECT="libprotoc 3.19.4"
if [[ "$VER" != "$EXPECT" ]]; then
  echo "'protoc' tool is version [$VER], expected [$EXPECT]." >&2
  exit 1
fi

GO_PLUGIN_PATH=$(command -v protoc-gen-gogo 2>&1)
if [[ "$GO_PLUGIN_PATH" == "" ]]; then
  echo "'protoc-gen-go' tool is required but missing." >&2
  exit 1
fi

go mod vendor

# global - change absolute path to relative
rm -rf tmp_php_proto && mkdir -p tmp_php_proto
cp -R chargehive tmp_php_proto
gsed -i'' -E 's#^import "github.com/chargehive/proto/chargehive/#import "chargehive/#g' tmp_php_proto/chargehive/**/*.proto

# go
rm -rf golang && mkdir -p golang
$PROTOC_PATH \
  -I ./vendor \
  -I ./tmp_php_proto \
  --gogo_out=plugins=grpc,paths=source_relative:golang \
  chargehive/**/*.proto

# php - remove gogo registration from php
gsed -i'' -E 's#^import "github.com/gogo/protobuf/gogoproto/gogo.proto";$##g' tmp_php_proto/chargehive/**/*.proto
gsed -i'' -E 's#^option \(gogoproto\..+##g' tmp_php_proto/chargehive/**/*.proto

rm -rf php && mkdir -p php
$PROTOC_PATH \
  -I ./vendor \
  -I ./tmp_php_proto \
  --php_out=php \
  tmp_php_proto/chargehive/**/*.proto

# cleanup
rm -Rf tmp_php_proto
