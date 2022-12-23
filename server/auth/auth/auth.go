package auth

import (
	"context"

	authpb "coolcar/auth/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ authpb.AuthServiceServer = &Service{}

type Service struct {
	authpb.UnimplementedAuthServiceServer
	OpenIDResolver
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
	return &authpb.LoginResponse{
		AccessToken: "X-OPENID:" + openID,
		ExpiresIn:   179875,
	}, nil
}
