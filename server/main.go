package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"coolcar/proto/gen/go"
	"coolcar/service/trip"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(l))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		":8081",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)

	if err != nil {
		log.Fatalf("can not start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("can not listen and serve: %v", err)
	}
}
