//go:generate ./generate.sh

package main

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/proto"
)

func main() {}
