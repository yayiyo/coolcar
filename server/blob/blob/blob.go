package blob

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"coolcar/blob/api/gen/v1"
	"coolcar/blob/dao"
	"coolcar/shared/id"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage interface {
	SignURL(ctx context.Context, method string, path string, timeout time.Duration) (url string, err error)
	Get(ctx context.Context, path string) (io.ReadCloser, error)
}

type Service struct {
	blobpb.UnimplementedBlobServiceServer
	Storage
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

func (s *Service) CreateBlob(ctx context.Context, req *blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	aid := id.AccountID(req.AccountId)
	br, err := s.Mongo.CreateBlob(ctx, aid)
	if err != nil {
		s.Logger.Error("create blob error", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	u, err := s.Storage.SignURL(ctx, http.MethodPut, br.Path, secToDuration(req.UploadUrlTimeoutSec))
	if err != nil {
		s.Logger.Error("sign URL error", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &blobpb.CreateBlobResponse{
		Id:        br.ID.Hex(),
		UploadUrl: u,
	}, nil
}
func (s *Service) GetBlob(ctx context.Context, req *blobpb.GetBlobRequest) (*blobpb.GetBlobResponse, error) {
	br, err := s.getBlobRecord(ctx, id.BlobID(req.Id))
	if err != nil {
		return nil, err
	}

	r, err := s.Storage.Get(ctx, br.Path)
	if err != nil {
		s.Logger.Error("get blob error", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "can not get blob", zap.Error(err))
	}
	defer r.Close()

	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read blob", zap.Error(err))
	}
	return &blobpb.GetBlobResponse{
		Data: d,
	}, nil
}
func (s *Service) GetBlobURL(ctx context.Context, req *blobpb.GetBlobURLRequest) (*blobpb.GetBlobURLResponse, error) {
	br, err := s.getBlobRecord(ctx, id.BlobID(req.Id))
	if err != nil {
		return nil, err
	}

	u, err := s.SignURL(ctx, http.MethodGet, br.Path, secToDuration(req.TimeoutSec))
	if err != nil {
		s.Logger.Error("GetBlobURL sign URL error", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &blobpb.GetBlobURLResponse{
		Url: u,
	}, nil
}

func (s *Service) getBlobRecord(ctx context.Context, bid id.BlobID) (*dao.BlobRecord, error) {
	br, err := s.Mongo.GetBlob(ctx, bid)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "")
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	return br, nil
}

func secToDuration(sec int32) time.Duration {
	return time.Duration(sec) * time.Second
}
