package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAijvnfFPE6oydG6WTFt4BlVccPIw5DcglBUIGVJn/8C9f8oXA
B7jzvKM4K0D3o3qVU7atsX/kMmSsyFEmb4g4SsGtpNROTIxBY87b2ZtKGdne2GkB
tgGcW9cimHfBVp5Ed6zZvANo6QuWee4zbFPhe8UGWX5slsZ9TgfPg3vosuqIC2OC
22hSjAAO64XtDDceMHlXOqC+fWHMYSjFGsw9QLnDEf4THD1ESm72gpPn2Wn8EPxt
G2wPZgcEPRixGt9Gd+RAX9CQl9ACAveSCCn37mEyTtDfiQ9/A0q0T48aQCd+2xR3
LeRTqds5/ZbrHo4JzVf2M8N0xlBidqyh7UBGZwIDAQABAoIBAHSwILZImF9OF4kw
+btB6KBCCmnqDHdfzjBkDaf4353Gv/2fffyG2ekQ9uc8Tk7HuaoS+Qhu5xjK9zeZ
mRJpMOuaFpSfSwE/HCu+gMNuXNz5ly2jZnhXX7//GQsPfDs8GpvTY3Fch4DB8WI3
+1ykaiDqrnN5TvzvzPLDkR5yyHnOSXqRW4/eQQ7xjeqwAwPMLqc9Ha3kM035aYzb
y3t/CaHPbVNXBHe3+pAb4j+SrtF+HDsAC5qudZNGgrBlUDl32dOJiDW6boBjUKNW
IHOTj4jw1RNx5m0CKzhbKDIpWg/Pt6L2X0TZX2qSFQxO/0dWD6TzS8JfXkndWbXq
lFV38pECgYEA/CYJI3y0aL7VlgaOCUN0TxpJEHT+NStt/8FP5Bapn3wbd2sCoMuE
EQ4HY4hxbJr6YCdoUkjjKLNqd+ABue2YB9X2LWJd7kyLgG4TCUU48vTOoFJCtvu8
o/UBZTMUJ8JQAlZ4IvDcNAnm8QGIwFOGg6vMnt8tLE+d8RWfrCxLjsUCgYEAjFhv
FZylkFtA8G3/LxxdQLPdkgJ+6KK3c0RbL73emyK+v0nuiW5A9+ShBIcvNuCICLU9
oiTwrzv48w+nZQ6IfFs5tzjHLS26DqkQm2lYO/l4Eskq5A/AFL+b0GCw5v3Y7sak
9nsEpHhTuTUx5vsX2rT62ajZBsHitY3cjj8i0zsCgYBUQDlJfD+jyDABwwrumXVh
gPzs3Mqb5XkJvgP9yHzA520eB8mHBLmKIU/iuBJ+IYKLYl9/Lw+H5/spNtYc1AC6
jYmGPJn6J+Vs8lq1/EU9GQN5mkLkdVTgy9q9f8W6SzkErvPRaP2K/cwGt2aELOSD
VoI2i4fCiI/ToFAL7XkJqQKBgGc4AzsZ4oqxEqnBvJShf9Q+dQ8V1tCyG8oi1A3g
zv+6Jh/59/4LoKyw8duqsQbjTClbYuEF2h6HNSlOsgaZEbikP2aJ07NeliCCJzyl
1ccGS1FFss3Y56Ra5/Xxpym8OPgkEN9WyqtL9AebIACJW5n20SeUD0nw/xQYKcZU
mQBRAoGAHAHzE7S9QdllGWT6bNA4R9UsYM9t3KRnVbU8NVWu9Z0WcwMYRqQ2R11r
0TnEh0cMdj20tkGBngdzX3xGW1TZ5C1/OeXUcaY5JHiTN6tstNYzEWsuNwGpDP29
NhCREJmFJ1as011tI/1A7NH0B9hkMLr59YkUmW3QznnYYLnmXoo=
-----END RSA PRIVATE KEY-----`

func TestJWTTokenGen_GenerateToken(t *testing.T) {
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("parse rsa private key error: %v", err)
	}

	g := NewJWTGenerator("coolcar/auth", privKey)
	g.nowFunc = func() time.Time {
		return time.Unix(1672135117688, 0)
	}

	token, err := g.GenerateToken("63a82c7d1205ece4c85b1ee7", 2*time.Hour)
	if err != nil {
		t.Errorf("generate token error: %v", err)
	}
	t.Logf("token: %s", token)
}
