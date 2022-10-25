package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

func Interceptor1(pubKey string) (us grpc.UnaryServerInterceptor, err error) {

	f, err := os.Open(pubKey)
	if err != nil {
		return nil, err
	}
	pubByte, err := ioutil.ReadAll(f)

	pem, err := jwt.ParseRSAPublicKeyFromPEM(pubByte)
	if err != nil {
		return nil, err
	}


	i := &interceptor1{
		pubKey: pem, //好像没用
		verifier: &token.JWTTokenVerify{
			PublicKey: pem,
		},
	}
	return i.HandleReq,nil
}

type tokenverifier1 interface {
	Verify(token string) (string,error)
}

type interceptor1 struct {
	pubKey *rsa.PublicKey //好像没用
	verifier tokenverifier1
}


//拦截器主要作用是做一些验证，或者参数规则校验的事情，在这里主要是拦截token，获取用户aid，位接下来使用
func (i *interceptor1) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){	//grpc.UnaryServerInfo一些其他的帮助信息

	//1、取出http请求头部的authorization
	tkn,err := tokenFromContext1(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated,"Unauthenticated")
	}

	//if i.pubKey != nil {
	//
	//}

	//使用继承了解token的验签工具进行验证并去除subject，也就是aid的值
	aid, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated,"token not valid: %v",err)
	}


	//ctx.Value("")
	//将aid放入待context中，并定义自己context中Key的格式并返回传递给接下来的操作
	return handler(ContextWithAccountID1(ctx,aid),req)
}

func tokenFromContext1(c context.Context) (string, error) {
	m,ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Error(codes.Unauthenticated,"")
	}

	tkn := ""
	for _, v := range m["authorization"] {
		if strings.HasPrefix(v, "Bearer") {
			tkn = v[len("Bearer "):]
		}
	}

	if tkn == "" {
		return "",status.Error(codes.Unauthenticated,"")
	}

	return tkn,nil
}

type accountIDKey1 struct {

}

func ContextWithAccountID1(c context.Context, aid string) context.Context {
	return context.WithValue(c,accountIDKey1{},aid)
}

func AccountIDFromContext1(c context.Context) (string, error) {
	v := c.Value(accountIDKey1{})
	aid,ok := v.(string)

	if !ok {
		return "",status.Error(codes.Unauthenticated,"")
	}
	return aid,nil
}