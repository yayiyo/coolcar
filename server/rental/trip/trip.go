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

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	accountID, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	s.Logger.Info("create trip ", zap.String("start", req.Start.String()), zap.String("accountID", accountID.String()))
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrip(context.Context, *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (s *Service) GetTrips(context.Context, *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrips not implemented")
}
func (s *Service) UpdateTrip(context.Context, *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
