package trip

import (
	"context"

	"coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedTripServiceServer
	Logger *zap.Logger
}

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	accountID, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	s.Logger.Info("create trip ", zap.String("start", req.Start), zap.String("accountID", accountID.String()))
	return nil, status.Error(codes.Unimplemented, "")
}
