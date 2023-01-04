package auth

import (
	"context"
	"io/ioutil"
	"strings"

	"coolcar/shared/auth/token"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorization = "authorization"
	bearerPrefix  = "Bearer "
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
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	accountID, err := i.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token verification failed: %v", err)
	}
	return handler(ContextWithAccountID(ctx, AccountID(accountID)), req)
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

// AccountID defines the account id type
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// ContextWithAccountID news a context with account ID
func ContextWithAccountID(ctx context.Context, accountID AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey{}, accountID)
}

// AccountIDFromContext returns the account ID in the context,
// it will return unauthenticated error if no account id is available
func AccountIDFromContext(ctx context.Context) (AccountID, error) {
	accountID, ok := ctx.Value(accountIDKey{}).(AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return accountID, nil
}
