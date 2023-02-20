package main

import (
	"context"
	"log"

	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/car"
	"coolcar/car/dao"
	"coolcar/car/mq/amqpclt"
	"coolcar/car/sim"
	"coolcar/shared/server"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Fatal("can not connect to amqp:", zap.Error(err))
	}

	publisher, err := amqpclt.NewPublisher(amqpConn, "coolcar")
	if err != nil {
		logger.Fatal("can not create publisher:", zap.Error(err))
	}

	carConn, err := grpc.Dial(":8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("can not connect to grpc(:8084)", zap.Error(err))
	}

	sub, err := amqpclt.NewSubscriber(amqpConn, "coolcar", logger)
	if err != nil {
		logger.Fatal("can not create subscriber:", zap.Error(err))
	}

	simController := &sim.Controller{
		CarService: carpb.NewCarServiceClient(carConn),
		Logger:     logger,
		Subscriber: sub,
	}

	go simController.RunSimulations(context.Background())

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "car",
		Addr: ":8084",
		RegisterFunc: func(s *grpc.Server) {
			carpb.RegisterCarServiceServer(s, &car.Service{
				Mongo:     dao.NewMongo(client.Database("coolcar")),
				Logger:    logger,
				Publisher: publisher,
			})
		},
		Logger: logger,
	}))
}
