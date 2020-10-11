package serializer

import (
	"golang_grpc_microservice/data"
	"golang_grpc_microservice/pb"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/stretchr/testify/require"
)

func TestProtoBufToBinarySerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/plant.bin"
	jsonFile := "../tmp/plant.json"

	plant1 := data.NewPlant()
	err := WriteProtoBufToBinaryFile(plant1, binaryFile)
	require.NoError(t, err)

	plant2 := &pb.Plant{}
	err = ReadProtoBufFromBinaryFile(binaryFile, plant2)
	require.NoError(t, err)
	require.True(t, proto.Equal(plant1, plant2))

	err = WriteProtoBufToJSONFile(plant1, jsonFile)
	require.NoError(t, err)

}
