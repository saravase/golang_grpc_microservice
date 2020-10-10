package sample

import (
	"golang_grpc_microservice/pb"
)

func NewPlant() *pb.Plant {
	return &pb.Plant{
		Id:   randomInt(1, 100),
		Name: randomPlantName(),
		Category: &pb.Category{
			Category: randomPlantCategory(),
		},
		Price:       randomFloat(50.0, 500.0),
		Avatar:      randomPlantImage(),
		Description: randomPlantDescription(),
	}
}
