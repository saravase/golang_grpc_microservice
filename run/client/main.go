package main

import (
	"context"
	"flag"
	"golang_grpc_microservice/data"
	"golang_grpc_microservice/pb"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "Server address")
	flag.Parse()
	log.Printf("Dial Server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't dial server : ", err)
	}

	plantClient := pb.NewPlantServiceClient(conn)

	plant := data.NewPlant()
	req := &pb.CreatePlantRequest{
		Plant: plant,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := plantClient.CreatePlant(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("Plant already Exists")
		} else {
			log.Fatal("Can't create plant: ", err)
		}
		return
	}

	log.Printf("Created plant with id: %s", res.Id)

}
