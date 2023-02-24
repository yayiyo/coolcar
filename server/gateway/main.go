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
	"github.com/namsral/flag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	addr        = flag.String("addr", ":8080", "address to connect")
	authAddr    = flag.String("auth_addr", ":8081", "auth address to connect")
	tripAddr    = flag.String("trip_addr", ":8082", "trip address to connect")
	profileAddr = flag.String("profile_addr", ":8082", "profile address to connect")
	carAddr     = flag.String("car_addr", ":8084", "car address to connect")
)

func main() {
	var (
		err error
	)
	flag.Parse()
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
			addr:         *authAddr,
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "trip",
			addr:         *tripAddr,
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:         "profile",
			addr:         *profileAddr,
			registerFunc: rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},
		{
			name:         "car",
			addr:         *carAddr,
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

	logger.Sugar().Infof("gateway starting at %s\n", addr)
	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Sugar().Fatalf("can not listen and serve: %v", err)
	}
}
