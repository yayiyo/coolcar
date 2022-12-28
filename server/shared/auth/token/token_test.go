package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 与生成token的 private key 成对儿 /auth/token/jwt_test.go
const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAijvnfFPE6oydG6WTFt4B
lVccPIw5DcglBUIGVJn/8C9f8oXAB7jzvKM4K0D3o3qVU7atsX/kMmSsyFEmb4g4
SsGtpNROTIxBY87b2ZtKGdne2GkBtgGcW9cimHfBVp5Ed6zZvANo6QuWee4zbFPh
e8UGWX5slsZ9TgfPg3vosuqIC2OC22hSjAAO64XtDDceMHlXOqC+fWHMYSjFGsw9
QLnDEf4THD1ESm72gpPn2Wn8EPxtG2wPZgcEPRixGt9Gd+RAX9CQl9ACAveSCCn3
7mEyTtDfiQ9/A0q0T48aQCd+2xR3LeRTqds5/ZbrHo4JzVf2M8N0xlBidqyh7UBG
ZwIDAQAB
-----END PUBLIC KEY-----`

func TestJWTTokenVerifier_Verify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("failed to parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		token   string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:  "valid token",
			token: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb29sY2FyL2F1dGgiLCJzdWIiOiI2M2E4MmM3ZDEyMDVlY2U0Yzg1YjFlZTciLCJleHAiOjE2NzIxMzUxMjQ4OTMsImlhdCI6MTY3MjEzNTExNzY5M30.CiosOCgdOPjrL7p5vOqeZ6PptXfmZBU2P5d6nfphuVxO9NTJZuiahGROvQKdPEPRZfxcDL34jRHylCMh5Kb9IoFbxenQ57WomvGNOMOrsP0kqKUqI7uZFtYkYTJ-0XErL13zYDRfnbauK_Pv1UgBrs0txNnhQgBtU-Fj-SR6vB1clQlSUVpsEJqI1pIfStBMt4xPOsMium0eIbHIpykA-ZEh1y0flbIH9HRhxz1LxjNwLYEppZQhvTBtEL2ufmwQdB-Ptn0p1d7mMgh47oEWue5qwAv8MRkzTdfYyrf6ncHl5kCCyd_65tdXeQZIM4PGjH4F09RM_ke_VWKFLKjfmA",
			now:   time.Unix(1672135117693, 0).Add(30 * time.Second),
			want:  "63a82c7d1205ece4c85b1ee7",
		},
		{
			name:    "invalid token",
			token:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb29sY2FyL2F1dGgiLCJzdWIiOiI2M2E4MmM3ZDEyMDVlY2U0Yzg1YjFlZTciLCJleHAiOjE2NzIxMzUxMjQ4ODgsImlhdCI6MTY3MjEzNTExNzY4OH0.OF0yhEbxwC6asF0ucUuEzi2ajWWaLbtZy1_VjZNqb5XCcyJikkbvcgtOS4hrdOgbephNL_7-hFn0_E0G78FnJq30tJAWECDvW_NOALktVcQjfXISU81ai2PzI52Afax5EqX7CPUdziXYO4o29-XtBh70MVpaZ5x8g05Jm7hFC8lA-mT85er45IY6-3v77Tfb_p0XJCdulquf3vOZv6LJiJKR8NAXs-TnhOeDWhJ4iK6oh4OMb8BjmQr7Mog0WyDpM9EOedBbplO2bz6W2o1jHvYZ4WqWsVYKJ1UMLUjfAJR5K9i1QL_h2jDzcKBHGF2pFrdGI25NkhRquZNWA5PVog",
			now:     time.Now(),
			wantErr: true,
		},
		{
			name:    "bad token",
			token:   "bad_token",
			wantErr: true,
		},
		{
			name:    "wrong signature",
			token:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb29sY2FyL2F1dGgiLCJzdWIiOiI2M2E4MmM3ZDEyMDVlY2U0Yzg1YjFlZTciLCJleHAiOjE2NzIxMzUxMjQ4ODgsImlhdCI6MTY3MjEzNTExNzY4OH0.epucrCtskoqLBfcBS0FbepPk8dxp1BBsqK3C5lDs643zxmSBbBeyGrFHVynXEFrqi_mZFzdZdSzfGKOu6yyEmTjYVXwcsMJ2FLcGu6BNsvjbnW5L5ZcrIEt7igxgeFLeGRdkpNL1se52EmeuweJyjiA0Tnra2nh_5-NXHn_47W-QaQJLwKOhxm5wrF1dr5_aZS2rZzPj7y6FNqB3yklfUvx0qE9yvnRazv7Rp2lQsovw5DdTEWnQG___TMzQzaHjoGjgitC1dfnY9Hx1_7p6TTT-pYtOU28gr3WWaQDbtlAvOTOSRlhkrHSkmLvFYiUdXBJRAWpHnWNaU-9T7mbgZg",
			now:     time.Unix(1672135117693, 0).Add(30 * time.Second),
			wantErr: true,
		},
	}

	for _, c := range cases {
		jwt.TimeFunc = func() time.Time {
			return c.now
		}
		t.Run(c.name, func(t *testing.T) {
			accountID, err := v.Verify(c.token)
			if !c.wantErr && err != nil {
				t.Errorf("[%s] no want error, bug got: %v", c.name, err)
			}

			if c.wantErr && err == nil {
				t.Errorf("[%s] want error, got none", c.name)
			}

			if accountID != c.want {
				t.Errorf("[%s] want: %q, but got: %q", c.name, c.want, accountID)
			}
		})
	}
}
