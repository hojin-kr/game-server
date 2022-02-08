#!/bin/bash
protoc --go_out=./cmd --go_opt=paths=source_relative \
    --go-grpc_out=./cmd --go-grpc_opt=paths=source_relative \
    proto/haru.proto