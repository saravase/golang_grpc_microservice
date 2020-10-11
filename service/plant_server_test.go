package service

import (
	"context"
	"golang_grpc_microservice/data"
	"golang_grpc_microservice/pb"
	"log"
	"testing"

	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
)

func TestServerCreatePlant(t *testing.T) {

	plantWithNoId := data.NewPlant()
	plantWithNoId.Id = ""

	plantWithInvalidId := data.NewPlant()
	plantWithInvalidId.Id = "invalid-uuid"

	plantWithDuplicateId := data.NewPlant()
	plantWithDuplicateStore := NewInMemoryPlantStore()
	err := plantWithDuplicateStore.Save(plantWithDuplicateId)
	require.Nil(t, err)

	testcases := []struct {
		caseName string
		plant    *pb.Plant
		store    PlantStore
		code     codes.Code
	}{
		{
			caseName: "plant_with_valid_id",
			plant:    data.NewPlant(),
			store:    NewInMemoryPlantStore(),
			code:     codes.OK,
		},
		{
			caseName: "plant_with_no-id",
			plant:    plantWithNoId,
			store:    NewInMemoryPlantStore(),
			code:     codes.OK,
		},
		{
			caseName: "plant_with_invalid_id",
			plant:    plantWithInvalidId,
			store:    NewInMemoryPlantStore(),
			code:     codes.InvalidArgument,
		},
		{
			caseName: "plant_with_duplicate_id",
			plant:    plantWithDuplicateId,
			store:    plantWithDuplicateStore,
			code:     codes.AlreadyExists,
		},
	}

	for index := range testcases {
		tc := testcases[index]
		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreatePlantRequest{
				Plant: tc.plant,
			}
			server := NewPlantServer(tc.store)
			res, err := server.CreatePlant(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.plant.Id) > 0 {
					require.Equal(t, tc.plant.Id, res.Id)
				}
			} else {
				log.Println(tc.caseName, "Error: ", err)
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
