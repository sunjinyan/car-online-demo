package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"os"
	"strings"
)
const (
	authorizationHeader = "authorization"
	bearerPrefix = "Bearer "
)

type verifyToken interface {
	Verify(token string) (string,error)
}


type authInterceptor struct {
	pubKey  *rsa.PublicKey
	//由于验证方式比较灵活，所以需要想到使用接口的思想
	verify  verifyToken
}

func InterceptorInfo(pubKey string) (grpc.UnaryServerInterceptor, error) {

	//读取文件，将pem内容转换为rsa的pubkey
	file, err := os.Open(pubKey)
	if err != nil {
		return nil, fmt.Errorf("open file error, error msg %v",err)
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file error, error msg %v",err)
	}

	pem, err := jwt.ParseRSAPublicKeyFromPEM(all)
	if err != nil {
		return nil, fmt.Errorf("read pem error, error msg %v",err)
	}


	i := &authInterceptor{
		pubKey: pem,
		verify: &token.JWTTokenVerify{PublicKey: pem},
	}

	return i.GetResInter, nil
}


func  (i *authInterceptor) GetResInter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	//获取token
	token, err := getToken(ctx)
	if  err != nil  {
		return "",fmt.Errorf("cannot get token from context")
	}
	//从token中获取aid,验证Verifytkoen
	aid, err := i.verify.Verify(token)
	if  err != nil  {
		return "",fmt.Errorf("cannot get aid from token")
	}
	//将aid传递到下一个函数
	return handler(createContextWithAid(ctx,id.AccountID(aid)),req)
}

//将aid强制类型转化
type aidKey struct {

}



func createContextWithAid(ctx context.Context, aid id.AccountID) context.Context {
	ctx  = context.WithValue(ctx,aidKey{},aid)
	return ctx
}

func GetAidFromContext(ctx context.Context) (id.AccountID,error) {
	aid := ctx.Value(aidKey{})

	id,ok := aid.(id.AccountID)
	if  !ok  {
		return "",fmt.Errorf("cannot get aid from context")
	}
	return id,nil
}

func getToken(c context.Context) (string,error){

	/*
	//metadata的形式是这种类型
	var md map[string][]string = make(map[string][]string)
	var mdSlice []string  = make([]string,10,20)
	mdSlice = append(mdSlice, "Bearer dalsjfqoji")
	md["authorization"] = mdSlice*/

	//["aaa"] = ["dasdas","dasdas"]

	ic, ok := metadata.FromIncomingContext(c)

	if ok {
		return "", fmt.Errorf("can not get metadata")
	}

	az,ok := ic[authorizationHeader]

	if !ok {
		//没有带认证信息，是否让通过？
		return "", fmt.Errorf("can not get metadata")
	}

	tkn := ""
	for _,v :=  range az{
		if strings.HasPrefix("v",bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}

	return tkn, nil
}