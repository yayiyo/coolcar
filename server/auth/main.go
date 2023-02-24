package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"github.com/golang-jwt/jwt/v4"
	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	addr      = flag.String("addr", ":8081", "address to connect")
	appID     = flag.String("app_id", "", "wechat app id, required")
	appSecret = flag.String("app_secret", "", "wechat app secret, required")
	mongoURL  = flag.String("mongo_url", "mongodb://localhost:27017", "mongodb url to connect")
)

func main() {
	flag.Parse()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	if *appID == "" {
		logger.Fatal("app_id can not be empty")
	}

	if *appSecret == "" {
		logger.Fatal("app secret can not be empty")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*mongoURL))
	if err != nil {
		logger.Fatal("can not connect to MongoDB:", zap.Error(err))
	}

	privKey, err := ioutil.ReadFile("auth/private.key")
	if err != nil {
		logger.Fatal("can not read private key file:", zap.Error(err))
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		logger.Fatal("can not parse private key file:", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "auth",
		Addr: *addr,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OpenIDResolver: &wechat.Service{
					AppID:     *appID,
					AppSecret: *appSecret,
				},
				TokenGenerator: token.NewJWTGenerator("coolcar/auth", key),
				TokenExpire:    2 * time.Hour,
				Mongo:          dao.NewMongo(client.Database("coolcar")),
				Logger:         logger,
			})
		},
		Logger: logger,
	}))
}
