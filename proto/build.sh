#!/bin/bash
protoc --go_out=./cmd --go_opt=paths=source_relative \
    --go-grpc_out=./cmd --go-grpc_opt=paths=source_relative \
    proto/haru.proto
protoc proto/haru.proto --swift_out=./proto --grpc-swift_out=./proto
# protoc --csharp_out=./proto/csharp \
#    --plugin=protoc-gen-csharp_grpc=/Users/hojin/Work/hojin/haru/proto/plugins/grpc_csharp_plugin \
#    --csharp_grpc_out=./proto/csharp \
#    proto/haru.proto
