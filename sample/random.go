package sample

import (
	"golang_grpc_microservice/pb"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomPlantCategory() pb.Category_Category {
	switch rand.Intn(4) {
	case 1:
		return pb.Category_Flower
	case 2:
		return pb.Category_Tree
	case 3:
		return pb.Category_Vegetable
	case 4:
		return pb.Category_Fancy
	default:
		return pb.Category_Fruit
	}
}

func randomInt(min int, max int) int32 {
	return int32(min + rand.Intn(max-min+1))
}

func randomPlantName() string {
	return randomStringFromSet(
		"Apple",
		"Orange",
		"Banana",
		"Neem",
		"Lemon",
	)
}

func randomFloat(min float32, max float32) float32 {
	return float32(min + rand.Float32() + (max - min))
}

func randomPlantImage() string {
	return randomStringFromSet(
		"apple.png",
		"orange.jpg",
		"banana.png",
		"neem.jpg",
		"lemon.png",
	)
}

func randomPlantDescription() string {
	return randomStringFromSet(
		"Sweetest Fruit",
		"High Vitamin Fruit",
		"Sweetest Fruit",
		"Powerful Tree",
		"Antibiotic Tree",
	)
}

func randomStringFromSet(strings ...string) string {
	i := len(strings)
	if i == 0 {
		return ""
	}
	return strings[rand.Intn(i)]
}
