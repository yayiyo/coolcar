package main

import (
	"log"
	"net"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("can not listen on :8081", zap.Error(err))
	}
	s := grpc.NewServer()
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})

	err = s.Serve(l)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
