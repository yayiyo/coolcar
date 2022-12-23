package main

import (
	"log"
	"net"

	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/wechat"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("can not listen on :8081", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx006574c1921658af",
			AppSecret: "8bde3a5eb25d40cd58501ed7e3dca226",
		},
		Logger: logger,
	})

	err = s.Serve(l)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
