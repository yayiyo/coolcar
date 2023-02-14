package profile

import (
	"context"
	"time"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile/dao"
	"coolcar/shared/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedProfileServiceServer
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

func (s *Service) GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("get profile", zap.String("aid", aid.String()))

	p, err := s.Mongo.GetProfile(ctx, aid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &rentalpb.Profile{}, nil
		}
		s.Logger.Error("can't get profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return p, nil
}

func (s *Service) SubmitProfile(ctx context.Context, i *rentalpb.Identity) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{
		Identity:       i,
		IdentityStatus: rentalpb.IdentityStatus_PENDING,
	}

	err = s.Mongo.UpdateProfile(ctx, aid, rentalpb.IdentityStatus_PENDING, p)
	if err != nil {
		s.Logger.Error("can't update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	go func() {
		time.Sleep(time.Second * 3)
		err = s.Mongo.UpdateProfile(context.Background(), aid, rentalpb.IdentityStatus_PENDING, &rentalpb.Profile{
			Identity:       i,
			IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
		})
		if err != nil {
			s.Logger.Error("verify the licence error", zap.Error(err))
		}
	}()

	return p, nil
}

func (s *Service) ClearProfile(ctx context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{}
	err = s.Mongo.UpdateProfile(ctx, aid, rentalpb.IdentityStatus_VERIFIED, p)
	if err != nil {
		s.Logger.Error("can't update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return p, nil
}
