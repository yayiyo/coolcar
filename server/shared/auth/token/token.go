package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// JWTTokenVerifier verify jwt token
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifies jwt token and returns account id
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("can't parse token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("token is not valid")
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", fmt.Errorf("claims is not a jwt.RegisteredClaims")
	}

	if err = claims.Valid(); err != nil {
		return "", fmt.Errorf("token is not valid: %v", err)
	}

	return claims.Subject, nil
}
