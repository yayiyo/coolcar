package main

import (
	"context"
	"log"

	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/blob"
	"coolcar/blob/cos"
	"coolcar/blob/dao"
	"coolcar/shared/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Fatal("can not connect to MongoDB:", zap.Error(err))
	}

	st, err := cos.NewService(
		"https://coolcar-1255667223.cos.ap-beijing.myqcloud.com",
		"AKIDwqiU9g5LRRM6h9jDVbT8e0AGKFxQhrpo",
		"sbFzZ5QaWPAi3A4kEf6NbuAKPL6rkdZW",
	)
	if err != nil {
		logger.Fatal("can not create cos service:", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "blob",
		Addr: ":8083",
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Mongo:   dao.NewMongo(client.Database("coolcar")),
				Logger:  logger,
				Storage: st,
			})
		},
		Logger: logger,
	}))
}
