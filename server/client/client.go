package main

import (
	"context"
	"fmt"
	"log"

	trippb "coolcar/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := trippb.NewTripServiceClient(conn)
	resp, err := c.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
