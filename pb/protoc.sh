#!/bin/sh
   protoc -I. \
     --go_out=plugins=grpc:. routeguide/routeguide.proto

   protoc -I$PROTOCPATH/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --go_out=plugins=grpc:. \
     trac.proto


