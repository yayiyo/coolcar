package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEApM+B1kPghEm8FMOXPGaWczkkgOmH+cSFxuWjfN2dq+czOLE6
o1QCBUZMLumhhDtVpga5rSEYcnCDGznqRRr8sgFgJuuP1F8fe9+0n1Cdv3lRpuR3
SeZK9TUpzoTDaSosS0bKiHB1pWcFFR9GanVP39S46rlqB4wJDbq58oORD4LshHAL
f6xuiC3Vz0S1v+NUMaZkhhJ/FdgfNj9zljYvOhdTdjE1JXJTOSo9fuGDHCyiMS9j
ir/V+9dO+rJvTqLHuMsmaP4TAm4vccr9ELsgybSH2EqFDL3EYRfOC/riu8nFyA4a
paUnrw4jPXdIxeSqFnXGVakf95zDyQwG0uU/YQIDAQABAoIBAQCZG1kgB1DFNIaw
t3+BEkDEbBp4U/sJUsOAogb+UhdPAmr4SNUTtFBoPAU8M7jj0gdgRLEroCpI7jMu
EOCNMGP+rf54SbAFeBUUjB8NPeQ+Y+Mx6t7S3UlXgPsEqxuqUA50JCC1Hdx9OzZi
h/pvneoFI2ZiULuqpTIn+gcSv5z8TYPX+VdOZzKkr+r/ev2a0HGmMv2gZxep1vcW
E2RXreZ8MDOtgob2dyyjQxFJ+DRWpCJk1kKFAuXmbN11tbO1ZNmUOb/ynz5988Dz
f9lJOkX7U19BtR/CgYG0PjUtroHDuEiH94KsFdwFlMlWDa6VnHXytM/FSbpFbKfA
QLwFD78BAoGBANG5kW9jkzDxOeg6ow+R5bXtPC2mzzQI2zPcevFVPt26ZzgaLOC2
CAW5Bz60HCN0+/4lJ95haluBoTG0lK21Adsm7JZFP485E31Jw0+CvN7qpJjWOpei
B5+gYsnJtg+NyWNwEdbo+eVAXRch5gIaajDkTMQQC32E2yXqYzO5hc1RAoGBAMks
7bmOlIdnuOpKcCMuK9okBL7D0+u0Cp5she1dvzAe7b7SXBk6R15yvL7EAN/QzG+3
8p3eHe7+ryU5BIRx/wPEbSWIu+P5cpVuskvk24Zfibodn3w5vc7Zm7O5nVease0N
6NX23pDbs5Q0ZYw8nnjyeKdn9VcASz87p5mQ3Y0RAoGBAL1p8fY+YpPTak9Zlifb
xzHWP4GjpIQEc6WVPdx090BeuBatXVbeUMSKZga1uKw1XdodSLLKHLrkiudPhvCU
CEccEpVtmYgLLpT7Z9CJ7XcPSPVYlraYenYq7s38xdeqYPbUIuiMphXtWaYo0YmY
vcvGhaaWLyqAMUU/ObVfm37hAoGAQdzDt4xGdE4w7AOS0vG6yaAhLZNPPkujblp5
Pk0C2u5FR8P3Awthcjp+MLZa0uu5Appmg+jERAp9rOIN6I6pvsIAOdmaKfjw2ptz
JAW5GTUOEjDAlhsRWTdFEoIiURwERGfZYrOACkzjbhH5bQAroc1AUw5l9CXUfM86
/7u9zQECgYB96ix7/l6GJvh2fFrSXcYWnNaGt9sAew12O/azJ5pRMwEaVC42UrMs
zS1NGDgO+P9JVtyhQ/KYf+a9q2BapaxHdVMRY+3YlYo1ySVcF1caXbaJdY04ekd7
Z9F889Rl3VLL1MMwn+etWtYKGdWRMO7xgc3QmOCtNQWwfX9iGVY/ZQ==
-----END RSA PRIVATE KEY-----`

func TestJWTTokenGen_GenerateToken(t *testing.T) {
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("parse rsa private key error: %v", err)
	}

	g := NewJWTTokenGen("coolcar/auth", privKey)
	g.nowFunc = func() time.Time {
		return time.Unix(1672135117693, 0)
	}

	token, err := g.GenerateToken("63a82c7d1205ece4c85b1ee7", 2*time.Hour)
	if err != nil {
		t.Errorf("generate token error: %v", err)
	}
	t.Logf("token: %s", token)
}
