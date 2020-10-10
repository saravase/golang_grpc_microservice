# golang_grpc_microservice

## Golang:

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency

## gRPC Environment Setup:

### Install Protocol Buffer Compiler and Plugin:
#### Linux:
    apt install -y protobuf-compiler
    protoc --version  # Ensure compiler version is 3+

### Plugin:
    go get -u github.com/golang/protobuf/protoc-gen-go

#### Get grpc package:
    go get google.golang.org/grpc

#### Initialize ptoto file


