package main

import (
	"context"
	"log"
	"net/http"
	"net/textproto"

	authpb "coolcar/auth/api/gen/v1"
	carpb "coolcar/car/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var (
		err error
	)
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create zap logger: %v", err)
	}

	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
			UseProtoNames:  true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{},
	}), runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		if key == textproto.CanonicalMIMEHeaderKey(runtime.MetadataHeaderPrefix+auth.ImpersonateAccountHeader) {
			return "", false
		}
		return runtime.DefaultHeaderMatcher(key)
	}))

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         ":8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "trip",
			addr:         ":8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:         "profile",
			addr:         ":8082",
			registerFunc: rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},
		{
			name:         "car",
			addr:         ":8084",
			registerFunc: carpb.RegisterCarServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		err = s.registerFunc(c, mux, s.addr, []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
		if err != nil {
			logger.Sugar().Fatalf("failed to register %s server: %v", s.name, err)
		}
	}

	addr := ":8080"
	logger.Sugar().Infof("gateway starting at %s\n", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Sugar().Fatalf("can not listen and serve: %v", err)
	}
}
