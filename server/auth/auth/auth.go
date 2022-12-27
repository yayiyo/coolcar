package auth

import (
	"context"
	"time"

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
	TokenGenerator
	TokenExpire time.Duration
	Mongo       *dao.Mongo
	Logger      *zap.Logger
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// TokenGenerator generates a token for the given account id
type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
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
		return nil, status.Error(codes.Internal, "")
	}

	token, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		s.Logger.Error("can't generate token", zap.String("accountID", accountID), zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
