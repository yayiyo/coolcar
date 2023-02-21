package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ImpersonateAccountHeader = "impersonate-account-id"
	authorization            = "authorization"
	bearerPrefix             = "Bearer "
)

func Interceptor(publicKey string) (grpc.UnaryServerInterceptor, error) {
	data, err := ioutil.ReadFile(publicKey)
	if err != nil {
		err = errors.Wrap(err, "failed to read public key file")
		return nil, err
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		err = errors.Wrap(err, "failed to parse public key")
		return nil, err
	}
	i := &interceptor{
		tokenVerifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}

	return i.HandleReq, nil
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	aid := impersonateFromContext(ctx)
	if aid != "" {
		fmt.Println("aid: ", aid)
		return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
	}
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	accountID, err := i.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token verification failed: %v", err)
	}
	return handler(ContextWithAccountID(ctx, id.AccountID(accountID)), req)
}

func impersonateFromContext(ctx context.Context) string {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}
	return imp[0]
}

func tokenFromContext(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}

	token := ""
	for _, v := range md[authorization] {
		if strings.HasPrefix(v, bearerPrefix) {
			token = v[len(bearerPrefix):]
			break
		}
	}

	if token == "" {
		return "", unauthenticated
	}
	return token, nil
}

type accountIDKey struct {
}

// ContextWithAccountID news a context with account ID
func ContextWithAccountID(ctx context.Context, accountID id.AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey{}, accountID)
}

// AccountIDFromContext returns the account ID in the context,
// it will return unauthenticated error if no account id is available
func AccountIDFromContext(ctx context.Context) (id.AccountID, error) {
	accountID, ok := ctx.Value(accountIDKey{}).(id.AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return accountID, nil
}
