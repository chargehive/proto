#!/bin/bash

mkdir genproto
cp proto/*.proto genproto/
sed -i '' "s#github.com/chargehive/proto/proto/##g" genproto/*.proto

protoc -I genproto \
-I $GOPATH/src \
--gogo_out=plugins=grpc,paths=source_relative:golang/chtype \
genproto/*.proto
#--php_out=plugins=grpc,paths=source_relative:php \

rm -Rf genproto