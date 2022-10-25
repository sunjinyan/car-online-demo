package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"time"
)

var Logger *zap.Logger

func InitLogger()  {
	Logger,err  := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("init logger failed,error message %v\n",err)
		return
	}
	Logger.Info("Init Logger Success")
}

func init() {
	InitLogger()
}

func main() {

	ctx,cancel :=  context.WithCancel(context.Background())
	defer cancel()
	//if err != nil {
	//	Logger.Fatal("listen port 8081 failed,error message ",zap.Error(err))
	//}

	//创建grpc gateway运行时服务，以及进行服务运行时选项配置
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,&runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{
			Multiline:         false,
			Indent:            "",
			AllowPartial:      false,
			UseProtoNames:     true,
			UseEnumNumbers:    true,
			EmitUnpopulated:   true,
			Resolver:          nil,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			AllowPartial:      false,
			DiscardUnknown:    false,
			Resolver:          nil,
		},
	}))


	/********************************************要学会规划代码************************************************************
	//注册各种服务
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), //使用http，非安全连接
	}) //如果链接，使用http铭文链接，不是用https)

	if err != nil{
		Logger.Sugar().Fatalf("grpc gateway err: v%",zap.Error(err))
	}

	err = rentalpb.RegisterTripServiceHandlerFromEndpoint(ctx, mux, "localhost:8082", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), //使用http，非安全连接
	}) //如果链接，使用http铭文链接，不是用https)

	if err != nil{
		Logger.Sugar().Fatalf("grpc gateway err: v%",zap.Error(err))
	}


	err = http.ListenAndServe(":8080", mux)

	if err != nil{
		Logger.Sugar().Fatalf("http service err: v%",zap.Error(err))
	}
	*******************************************要学会规划代码*************************************************************/
	//统一规划注册所有服务到gateway
	var RegisterServiceHandler = []struct {
		endpoint    string
		name        string
		registerFun func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			endpoint: "localhost:8081",
			name: "auth",
			registerFun: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			endpoint: "localhost:8082",
			name: "trip",
			registerFun: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	dailOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), //使用http，非安全连接
	}

	for _,hand := range RegisterServiceHandler {
		err := hand.registerFun(ctx, mux, hand.endpoint, dailOptions)
		if err != nil{
			Logger.Sugar().Fatalf("http service err: v%",zap.String("name",hand.name),zap.Error(err))
		}
	}
	Logger.Sugar().Fatalf("http service err: v%",zap.Error(http.ListenAndServe(":8080", mux)))
}

func JWTCrypt(aid string,expire time.Duration,privateKey *rsa.PrivateKey) {
	claims := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: 0,
		IssuedAt:  0,
		Issuer:    "",
		NotBefore: 0,
		Subject:   aid,
	})

	token, _ := claims.SignedString(privateKey)
	fmt.Println(token)
}

func JWTDecrype(key []byte,token string) string {
	pem, _ := jwt.ParseRSAPublicKeyFromPEM(key)

	parse, _ := jwt.ParseWithClaims(token,jwt.StandardClaims{}, func(token2 *jwt.Token) (interface{}, error) {
		return pem, nil
	})

	if parse.Valid {
		return ""
	}

	if err := parse.Claims.Valid(); err != nil {
		return ""
	}

	claims,ok := parse.Claims.(*jwt.StandardClaims)
	if !ok  {

	}
	return claims.Subject
}