package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
)

type JWTTokenVerify struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerify)Verify(token string) (string,error)  {
	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.PublicKey, nil //第三个参数是个KeyFunc，并不是把PublicKey当作一个参数传过来，而是作为一个函数，为什么呢？是因为我们的token使用的是什么算法都不知道，只有见到token第一段才能解出来RS512的算法
		//所以说不能直接给一个PublicKey，需要将第一段密文使用base64解出来，确定其中的alg是什么，根据该值来进行不同的PublicKey返回，
		//因为加密那边可能有的时候使用的是RS512有的时候是RS256或者其他的，那么所使用的PublicKey就会不同，如果只使用一种，那么直接返回PublicKey即可
	})//向jwt.StandardClaims写数据，也就是解出来的数据放在这里，这里和加密的时候一致就好
	if err != nil {
		return "", fmt.Errorf("cannot parse token %v",err)
	}
	if !claims.Valid {
		return "", fmt.Errorf("cannot Valid token %v",err)
	}

	clm,ok := claims.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("cannot claims  token %v",err)
	}

	if err = clm.Valid();err != nil {
		return "", fmt.Errorf("cannot claims Valid token %v",err)
	}

	return clm.Subject,nil
}
