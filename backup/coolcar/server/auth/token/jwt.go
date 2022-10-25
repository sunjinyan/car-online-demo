package token

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"time"
)

//实现者决定使用JWT来实现Token的生成接口
type JWTTokenGenerate struct {
	issuer string
	nowFunc func() time.Time
	privateKey *rsa.PrivateKey
}

func NewJWTTokenGenerate(issure string,privateKey *rsa.PrivateKey) *JWTTokenGenerate  {
	return &JWTTokenGenerate{
		issuer:  issure,
		nowFunc: func() time.Time {
			return time.Now()
		},
		privateKey: privateKey,
	}
}

func (j *JWTTokenGenerate)GenerateToken(accountID string,expire time.Duration)(string,error)  {
	nowSec := j.nowFunc().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		//Audience:  "",
		ExpiresAt: nowSec + int64(expire.Seconds()),
		//Id:        "",
		IssuedAt:  time.Now().Unix(),
		Issuer:    j.issuer,
		NotBefore: 0,
		Subject:   accountID,
	})
	//{
	//	"sub": "630734ab70541b5f24913914",
	//	"name": "John Doe",
	//	"iat": 1516239022,
	//	"exp": 1516315000,
	//	"nbf": 60,
	//	"iss": "coolcar/auth"
	//}
	return token.SignedString(j.privateKey)
}