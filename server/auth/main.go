package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"time"

	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"github.com/golang-jwt/jwt/v4"
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

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("can not listen on :8081", zap.Error(err))
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Fatal("can not connect to MongoDB:", zap.Error(err))
	}

	s := grpc.NewServer()
	privKey, err := ioutil.ReadFile("auth/private.key")
	if err != nil {
		logger.Fatal("can not read private key file:", zap.Error(err))
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		logger.Fatal("can not parse private key file:", zap.Error(err))
	}

	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx006574c1921658af",
			AppSecret: "8bde3a5eb25d40cd58501ed7e3dca226",
		},
		TokenGenerator: token.NewJWTTokenGen("coolcar/auth", key),
		TokenExpire:    2 * time.Hour,
		Mongo:          dao.NewMongo(client.Database("coolcar")),
		Logger:         logger,
	})

	err = s.Serve(l)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
