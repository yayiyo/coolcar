package main

import (
	"context"
	"log"

	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/blob"
	"coolcar/blob/cos"
	"coolcar/blob/dao"
	"coolcar/shared/server"
	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	addr      = flag.String("addr", ":8083", "address to connect")
	cosAddr   = flag.String("cos_url", "", "cos address to connect")
	secretID  = flag.String("secret_id", "", "cos secret id, required")
	secretKey = flag.String("secret_key", "", "cos secret key, required")
	mongoURL  = flag.String("mongo_url", "mongodb://localhost:27017", "mongodb url to connect")
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	if *cosAddr == "" {
		logger.Fatal("cos address cannot be empty")
	}

	if *secretID == "" {
		logger.Fatal("secret_id can not be empty")
	}

	if *secretKey == "" {
		logger.Fatal("secret_key can not be empty")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*mongoURL))
	if err != nil {
		logger.Fatal("can not connect to MongoDB:", zap.Error(err))
	}

	st, err := cos.NewService(
		*cosAddr,
		*secretID,
		*secretKey,
	)
	if err != nil {
		logger.Fatal("can not create cos service:", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "blob",
		Addr: *addr,
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
