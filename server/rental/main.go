package main

import (
	"context"
	"log"
	"time"

	blobpb "coolcar/blob/api/gen/v1"
	carpb "coolcar/car/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile"
	profileDao "coolcar/rental/profile/dao"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	profileClient "coolcar/rental/trip/client/profile"
	tripDao "coolcar/rental/trip/dao"
	"coolcar/shared/server"
	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr     = flag.String("addr", ":8082", "address to connect")
	blobAddr = flag.String("bolb_addr", ":8083", "blob address to connect")
	carAddr  = flag.String("car_addr", ":8084", "car address to connect")
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

	conn, err := grpc.Dial(*blobAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("can not connect to grpc(:8083):", zap.Error(err))
	}

	profileService := &profile.Service{
		BlobService:       blobpb.NewBlobServiceClient(conn),
		PhotoGetExpire:    5 * time.Second,
		PhotoUploadExpire: 10 * time.Second,
		Mongo:             profileDao.NewMongo(client.Database("coolcar")),
		Logger:            logger,
	}

	carConn, err := grpc.Dial(*carAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("can not connect to grpc(:8084):", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              *addr,
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				ProfileManager: &profileClient.Manager{
					Fetcher: profileService,
				},
				POIManager: &poi.Manager{},
				CarManager: &car.Manager{
					CarService: carpb.NewCarServiceClient(carConn),
				},
				Mongo:  tripDao.NewMongo(client.Database("coolcar")),
				Logger: logger,
			})
			rentalpb.RegisterProfileServiceServer(s, profileService)
		},
		Logger: logger,
	}))
}
