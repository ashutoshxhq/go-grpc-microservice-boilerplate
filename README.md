# go-grpc-microservice-boilerplate
A boilerplate using Go and gRPC

# Setup
- Make sure `protoc` is installed and available in path
- Make sure you have golang properly setup
- Make sure you have grpc-gateway properly setup

## grpc-gateway setup
```
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

# Todo
- Application Structure
- Logging
- Debugging
- Benchmarking
- Profiling
- Documenting
