package profile

import (
	"context"
	"time"

	blobpb "coolcar/blob/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedProfileServiceServer
	BlobService       blobpb.BlobServiceClient
	PhotoGetExpire    time.Duration
	PhotoUploadExpire time.Duration
	Mongo             *dao.Mongo
	Logger            *zap.Logger
}

func (s *Service) GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("get profile", zap.String("aid", aid.String()))

	p, err := s.Mongo.GetProfile(ctx, aid)
	if err != nil {
		code := s.logAndConvertProfileErr(err)
		if code == codes.NotFound {
			return &rentalpb.Profile{}, nil
		}
		return nil, status.Error(code, "")
	}

	if p.Profile == nil {
		return &rentalpb.Profile{}, nil
	}
	return p.Profile, nil
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

	err = s.Mongo.UpdateProfile(ctx, aid, rentalpb.IdentityStatus_NOT_SUBMITTED, p)
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

func (s *Service) GetProfilePhoto(ctx context.Context, req *rentalpb.GetProfilePhotoRequest) (*rentalpb.GetProfilePhotoResponse, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	p, err := s.Mongo.GetProfile(ctx, aid)
	if err != nil {
		return nil, status.Error(s.logAndConvertProfileErr(err), "")
	}

	if p.PhotoBlobID == "" {
		return nil, status.Error(codes.NotFound, "")
	}

	br, err := s.BlobService.GetBlobURL(ctx, &blobpb.GetBlobURLRequest{
		Id:         p.PhotoBlobID,
		TimeoutSec: int32(s.PhotoGetExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("get blob url error", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &rentalpb.GetProfilePhotoResponse{
		Url: br.Url,
	}, nil
}
func (s *Service) CreateProfilePhoto(ctx context.Context, req *rentalpb.CreateProfilePhotoRequest) (*rentalpb.CreateProfilePhotoResponse, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	br, err := s.BlobService.CreateBlob(ctx, &blobpb.CreateBlobRequest{
		AccountId:           aid.String(),
		UploadUrlTimeoutSec: int32(s.PhotoUploadExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("create blob error", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}

	err = s.Mongo.UpdateProfilePhoto(ctx, aid, id.BlobID(br.Id))
	if err != nil {
		s.Logger.Error("update profile photo error", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}

	return &rentalpb.CreateProfilePhotoResponse{
		UploadUrl: br.UploadUrl,
	}, nil
}
func (s *Service) CompleteProfilePhoto(ctx context.Context, req *rentalpb.CompleteProfilePhotoRequest) (*rentalpb.Identity, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	p, err := s.Mongo.GetProfile(ctx, aid)
	if err != nil {
		return nil, status.Error(s.logAndConvertProfileErr(err), "")
	}

	if p.PhotoBlobID == "" {
		return nil, status.Error(codes.NotFound, "")
	}

	blob, err := s.BlobService.GetBlob(ctx, &blobpb.GetBlobRequest{
		Id: p.PhotoBlobID,
	})
	if err != nil {
		s.Logger.Error("failed to get profile blob", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}

	// TODO：OCR
	s.Logger.Info("profile photo blob", zap.Int("size", len(blob.Data)))

	return &rentalpb.Identity{
		LicNumber: "88888888888888",
		Name:      "李菜鸡",
		Gender:    rentalpb.Gender_MALE,
		BirthDate: "1998-02-03",
	}, nil
}

func (s *Service) ClearProfilePhoto(ctx context.Context, req *rentalpb.ClearProfilePhotoRequest) (*rentalpb.ClearProfilePhotoResponse, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = s.Mongo.UpdateProfilePhoto(ctx, aid, "")
	if err != nil {
		s.Logger.Error("failed to clear profile photo", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &rentalpb.ClearProfilePhotoResponse{}, nil
}

func (s *Service) logAndConvertProfileErr(err error) codes.Code {
	if err == mongo.ErrNoDocuments {
		return codes.NotFound
	}
	s.Logger.Error("can't get profile", zap.Error(err))
	return codes.Internal
}
