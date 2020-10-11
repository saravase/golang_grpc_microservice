package service

import (
	"context"
	"errors"
	"golang_grpc_microservice/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PlantServer struct {
	Store PlantStore
}

func NewPlantServer(store PlantStore) *PlantServer {
	return &PlantServer{store}
}

func (server *PlantServer) CreatePlant(c context.Context, req *pb.CreatePlantRequest) (*pb.CreatePlantResponse, error) {
	plant := req.GetPlant()
	log.Printf("[DEBUG] Create plant request, plant data id: %s", plant.Id)

	if len(plant.Id) > 0 {
		_, err := uuid.Parse(plant.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "[ERROR] While, Validating plant id: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "[ERROR] While, Generating plant id: %v", err)
		}
		plant.Id = id.String()
	}

	if c.Err() == context.Canceled {
		log.Print("Client request cancelled")
		return nil, status.Error(codes.Canceled, "[ERROR] While, Cancelling client request")
	}

	if c.Err() == context.DeadlineExceeded {
		log.Print("Deadline exceed")
		return nil, status.Error(codes.DeadlineExceeded, "[ERROR] While, Time Exceed")
	}

	err := server.Store.Save(plant)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "[ERROR] While, Saving plant data: %v", err)
	}

	res := &pb.CreatePlantResponse{
		Id: plant.Id,
	}
	log.Printf("[DEBUG] Created plant with id: %s", plant.Id)
	return res, nil
}
