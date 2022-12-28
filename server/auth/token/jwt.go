package token

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTGenerator generates a JWT token
type JWTGenerator struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

func NewJWTGenerator(issuer string, privateKey *rsa.PrivateKey) *JWTGenerator {
	return &JWTGenerator{
		privateKey: privateKey,
		issuer:     issuer,
		nowFunc:    time.Now,
	}
}

// GenerateToken generates a token
func (j *JWTGenerator) GenerateToken(accountID string, expire time.Duration) (string, error) {
	now := j.nowFunc()
	t := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    j.issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expire)),
		Subject:   accountID,
	})
	return t.SignedString(j.privateKey)
}
