package service

import (
	"context"
	"golang_grpc_microservice/data"
	"golang_grpc_microservice/pb"
	"golang_grpc_microservice/serializer"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreatePlant(t *testing.T) {
	t.Parallel()

	plantServer, serverAddress := startTestPlantServer(t)
	plantClient := newTestPlantClient(t, serverAddress)

	plant := data.NewPlant()
	sourceId := plant.Id
	req := &pb.CreatePlantRequest{
		Plant: plant,
	}

	res, err := plantClient.CreatePlant(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, sourceId, res.Id)

	// check that the plant is saved to the store
	plantData, err := plantServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, plantData)

	// check that saved plant data is the same as one plant data we send
	validateSamePlant(t, plant, plantData)
}

func startTestPlantServer(t *testing.T) (*PlantServer, string) {
	plantServer := NewPlantServer(NewInMemoryPlantStore())

	grpcServer := grpc.NewServer()
	pb.RegisterPlantServiceServer(grpcServer, plantServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return plantServer, listener.Addr().String()
}

func newTestPlantClient(t *testing.T, serverAddress string) pb.PlantServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewPlantServiceClient(conn)
}

func validateSamePlant(t *testing.T, plant1 *pb.Plant, plant2 *pb.Plant) {
	json1, err := serializer.ProtoBufToJSON(plant1)
	require.NoError(t, err)

	json2, err := serializer.ProtoBufToJSON(plant2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
