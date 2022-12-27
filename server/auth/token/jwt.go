package token

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTTokenGen generates a JWT token
type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		privateKey: privateKey,
		issuer:     issuer,
		nowFunc:    time.Now,
	}
}

// GenerateToken generates a token
func (j *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	now := j.nowFunc()
	t := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    j.issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expire)),
		Subject:   accountID,
	})
	return t.SignedString(j.privateKey)
}
