#!/bin/bash

protoc -I proto \
-I $GOPATH/src \
--gogo_out=plugins=grpc,paths=source_relative:golang/chtype \
--proto_path=$GOPATH/src/github.com/chargehive/proto \
proto/*.proto
#--php_out=plugins=grpc,paths=source_relative:php \
