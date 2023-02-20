package main

import (
	"context"
	"fmt"
	"math/rand"

	carpb "coolcar/car/api/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := carpb.NewCarServiceClient(conn)
	for i := 0; i < 8; i++ {
		la := rand.Int31n(10)
		car, err := client.CreateCar(context.Background(), &carpb.CreateCarRequest{
			Car: &carpb.Car{
				Status: carpb.CarStatus_LOCKED,
				Position: &carpb.Location{
					Latitude:  float64(30 + la),
					Longitude: float64(120 + la),
				},
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Car %q is created\n", car.Id)
	}
}
