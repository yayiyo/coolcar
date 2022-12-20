package trip

import (
	"context"

	"coolcar/proto/gen/go"
)

var _ trippb.TripServiceServer = &Service{}

type Service struct {
	trippb.UnimplementedTripServiceServer
}

// GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
func (*Service) GetTrip(ctx context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "北京",
			End:         "上海",
			DurationSec: 3600,
			FeeCent:     8647,
			StartPos: &trippb.Location{
				Latitude:  31,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 123,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  33,
					Longitude: 123,
				},
				{
					Latitude:  36,
					Longitude: 134,
				},
			},
			Status: trippb.TripStatus_IN_PROGRESS,
		},
	}, nil
}
