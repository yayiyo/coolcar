package main

import (
	"context"
	"log"
	"net/http"

	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/car"
	"coolcar/car/dao"
	"coolcar/car/mq/amqpclt"
	"coolcar/car/sim"
	"coolcar/car/trip"
	"coolcar/car/ws"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"github.com/gorilla/websocket"
	"github.com/namsral/flag"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr     = flag.String("addr", ":8084", "address to connect")
	wsAddr   = flag.String("ws_addr", ":9090", "ws address to connect")
	tripAddr = flag.String("trip_addr", ":8082", "trip address to connect")
	mqAddr   = flag.String("mq_url", "amqp://guest:guest@localhost:5672/", "rabbitMQ address to connect")
	mongoURL = flag.String("mongo_url", "mongodb://localhost:27017", "mongodb url to connect")
)

func main() {
	flag.Parse()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*mongoURL))
	if err != nil {
		logger.Fatal("can not connect to MongoDB:", zap.Error(err))
	}

	amqpConn, err := amqp.Dial(*mqAddr)
	if err != nil {
		logger.Fatal("can not connect to amqp:", zap.Error(err))
	}

	publisher, err := amqpclt.NewPublisher(amqpConn, "coolcar")
	if err != nil {
		logger.Fatal("can not create publisher:", zap.Error(err))
	}

	carConn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	u := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	http.HandleFunc("/ws", ws.Handler(u, sub, logger))
	go func() {
		addr := *wsAddr
		logger.Info("starting http server", zap.String("at", addr))
		logger.Sugar().Fatal(http.ListenAndServe(addr, nil))
	}()

	// start trip updater
	tripConn, err := grpc.Dial(*tripAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect trip grpc :8082", zap.Error(err))
	}
	go trip.RunUpdater(sub, rentalpb.NewTripServiceClient(tripConn), logger)

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "car",
		Addr: *addr,
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
