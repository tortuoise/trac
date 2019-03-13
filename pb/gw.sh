#!/bin/sh
   protoc -I$PROTOCPATH/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --grpc-gateway_out=logtostderr=true:. \
     ./trac.proto


