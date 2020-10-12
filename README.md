# golang_grpc_unary_stream

## Golang:
    Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language. Programs are assembled by using packages, for efficient management of dependencies. This language also supports environment adopting patterns alike to dynamic languages

## gRPC:
    gRPC is a modern open source high performance RPC framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

## gRPC Environment Setup

#### Install Protocol Buffer Compiler:
    apt install -y protobuf-compiler
    protoc --version  # Ensure compiler version is 3+

#### Install Protocol Buffer Generate Plugin:
    go get -u github.com/golang/protobuf/protoc-gen-go

#### Get gRPC package:
    go get google.golang.org/grpc

#### Make proto file:
    proto/sample.proto

#### Initialize protobuf generate make command:
    gen:
	protoc --proto_path=proto/ proto/*.proto --go_out=plugins=grpc:pb


## Serialization with unit test:
    1. Protocol buffer message  =>  Binary data
    2. Protocol buffer message  =>  JSON data
    3. Binary data              =>  Protocol buffer message

##  External Packages:
    1. github.com/google/uuid
    2. github.com/stretchr/testify
    3. github.com/jinzhu/copier