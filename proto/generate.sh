#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
# cd `dirname $0`

function gen() {
  proto=$1
  proto_path="./java/src/main/proto/"
  proto_file=$proto_path/$proto.proto
  echo $proto
  protoc \
    --proto_path=$proto_path \
    --go_out=. \
    --go-grpc_out=. \
    $proto_file
}

gen hello
