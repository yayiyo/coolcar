package server

import (
	"net"

	"coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(gc *GRPCConfig) error {
	nameField := zap.String("name", gc.Name)
	l, err := net.Listen("tcp", gc.Addr)
	if err != nil {
		gc.Logger.Fatal("can not listen", nameField, zap.String("addr", gc.Addr), zap.Error(err))
	}

	var opts []grpc.ServerOption
	if gc.AuthPublicKeyFile != "" {
		var in grpc.UnaryServerInterceptor
		in, err = auth.Interceptor(gc.AuthPublicKeyFile)
		if err != nil {
			gc.Logger.Fatal("can not create auth interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}
	s := grpc.NewServer(opts...)
	gc.RegisterFunc(s)

	gc.Logger.Sugar().Infof("starting server %s at %s\n", gc.Name, gc.Addr)
	return s.Serve(l)
}
