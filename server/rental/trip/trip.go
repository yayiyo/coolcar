package trip

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedTripServiceServer
	ProfileManager
	CarManager
	POIManager
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

// ProfileManager defines the ACL (Anti Corruption Layer) for profile verification logic.
type ProfileManager interface {
	Verify(ctx context.Context, aid id.AccountID) (id.IdentityID, error)
}

// CarManager defines the ACL for car management.
type CarManager interface {
	Verify(ctx context.Context, carID id.CarID, location *rentalpb.Location) error
	Unlock(ctx context.Context, carID id.CarID) error
}

// POIManager resolves the POI(Point Of Interest).
type POIManager interface {
	Resolve(ct context.Context, location *rentalpb.Location) (string, error)
}

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	accountID, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// 验证驾驶者身份
	iID, err := s.ProfileManager.Verify(ctx, accountID)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	if req.CarId == "" || req.Start == nil {
		return nil, status.Error(codes.InvalidArgument, "参数不能为空")
	}

	// 检查车辆状态
	carID := id.CarID(req.CarId)
	err = s.CarManager.Verify(ctx, carID, req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	ls := s.calcCurrentStatus(ctx, &rentalpb.LocationStatus{
		Location:     req.Start,
		TimestampSec: nowFunc(),
	}, req.Start)
	// 创建行程：写入数据库，开始计费
	trip, err := s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId:  accountID.String(),
		CarId:      carID.String(),
		IdentityId: iID.String(),
		Start:      ls,
		Current:    ls,
		Status:     rentalpb.TripStatus_IN_PROGRESS,
	})
	if err != nil {
		s.Logger.Warn("can't create trip", zap.Error(err))
		return nil, status.Errorf(codes.AlreadyExists, "")
	}
	// 车辆开锁
	go func() {
		err = s.CarManager.Unlock(ctx, carID)
		if err != nil {
			s.Logger.Error("can't unlock car", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   trip.ID.Hex(),
		Trip: trip.Trip,
	}, nil
}

func (s *Service) GetTrip(ctx context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	accountID, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tr, err := s.Mongo.GetTrip(ctx, id.TripID(req.Id), accountID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return tr.Trip, nil
}
func (s *Service) GetTrips(ctx context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	accountID, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	trips, err := s.Mongo.GetTrips(ctx, accountID, req.Status)
	if err != nil {
		s.Logger.Error("can not get trips", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	res := &rentalpb.GetTripsResponse{
		Trips: make([]*rentalpb.TripEntity, 0),
	}

	for _, trip := range trips {
		res.Trips = append(res.Trips, &rentalpb.TripEntity{
			Id:   trip.ID.Hex(),
			Trip: trip.Trip,
		})
	}
	return res, nil
}
func (s *Service) UpdateTrip(ctx context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	trip, err := s.Mongo.GetTrip(ctx, id.TripID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}

	if trip.Trip.Current == nil {
		s.Logger.Error(fmt.Sprintf("the trip %q without current location status", req.Id))
		return nil, status.Error(codes.Internal, "")
	}

	cur := trip.Trip.Current
	if req.Current != nil {
		cur = req.Current
	}
	trip.Trip.Current = s.calcCurrentStatus(ctx, trip.Trip.Current, cur.Location)

	if req.Current != nil {
		trip.Trip.Current = s.calcCurrentStatus(ctx, trip.Trip.Current, req.Current.Location)
	}
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}

var nowFunc = func() int64 {
	return time.Now().Unix()
}

const (
	centsPerSec = 0.7
	kmPerSec    = 0.02
)

func (s *Service) calcCurrentStatus(ctx context.Context, last *rentalpb.LocationStatus, current *rentalpb.Location) *rentalpb.LocationStatus {
	now := nowFunc()
	elapsedSec := float64(now - last.TimestampSec)
	poi, err := s.POIManager.Resolve(ctx, current)
	if err != nil {
		s.Logger.Error("calcCurrentStatus can't resolve poi", zap.Stringer("location", current), zap.Error(err))
	}
	return &rentalpb.LocationStatus{
		Location:     current,
		FeeCent:      last.FeeCent + int32(centsPerSec*elapsedSec*2*rand.Float64()),
		KmDriven:     last.KmDriven + kmPerSec*elapsedSec*2*rand.Float64(),
		TimestampSec: now,
		PoiName:      poi,
	}
}
