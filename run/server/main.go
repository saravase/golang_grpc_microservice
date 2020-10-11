package main

import (
	"flag"
	"fmt"
	"golang_grpc_microservice/pb"
	"golang_grpc_microservice/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "Server port")
	flag.Parse()
	log.Printf("Start the server on port: %d", *port)

	plantServer := service.NewPlantServer(service.NewInMemoryPlantStore())
	grpcServer := grpc.NewServer()
	pb.RegisterPlantServiceServer(grpcServer, plantServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Can't start the server:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Can't start the server", err)
	}
}
