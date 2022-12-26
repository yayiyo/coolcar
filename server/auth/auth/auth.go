package auth

import (
	"context"

	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ authpb.AuthServiceServer = &Service{}

type Service struct {
	authpb.UnimplementedAuthServiceServer
	OpenIDResolver
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	openID, err := s.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "can't resolve OpenID: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(ctx, openID)
	if err != nil {
		s.Logger.Error("can't resolve OpenID %v error: %+v", zap.String("openid", openID), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: "X-ACCOUNT:" + accountID,
		ExpiresIn:   7200,
	}, nil
}
