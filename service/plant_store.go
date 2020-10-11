package service

import (
	"errors"
	"fmt"
	"golang_grpc_microservice/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("[ERROR] Plant record already exist")

type PlantStore interface {
	Save(plant *pb.Plant) error
	Find(id string) (*pb.Plant, error)
}

type InMemoryPlantStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Plant
}

func NewInMemoryPlantStore() *InMemoryPlantStore {
	return &InMemoryPlantStore{
		data: make(map[string]*pb.Plant),
	}
}

func (store *InMemoryPlantStore) Save(plant *pb.Plant) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[plant.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Plant{}
	err := copier.Copy(other, plant)
	if err != nil {
		return fmt.Errorf("[ERROR] While, Copying plant data: %v", err)
	}
	store.data[other.Id] = other
	return nil

}

func (store *InMemoryPlantStore) Find(id string) (*pb.Plant, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	plant := store.data[id]
	if plant == nil {
		return nil, nil
	}

	other := &pb.Plant{}
	err := copier.Copy(other, plant)
	if err != nil {
		return nil, fmt.Errorf("[ERRORO] While, Copying plant data: %v", err)
	}

	return other, nil
}
