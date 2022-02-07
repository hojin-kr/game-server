#!/bin/bash
protoc --go_out=./cmd/proto --go_opt=paths=source_relative \
    --go-grpc_out=./cmd/proto --go-grpc_opt=paths=source_relative \
    proto/haru.proto