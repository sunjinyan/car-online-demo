package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger() (logger *zap.Logger,err error) {
	//logger,err  = zap.NewDevelopment()
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
	//if err != nil {
	//	 log.Fatalf("init logger failed,error message %v\n",err)
	//	return
	//}
	//Logger.Info("Init Logger Success")
}

//func init() {
//	InitLogger()
//}
var (
	privateKey []byte
	publicKey []byte
)
func InitJwtKey() {
	file, err := os.Open("auth/private.key")
	if err != nil {
		log.Fatalf("open jwt private key file failed, error message: %v",err)
	}
	privateKey, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("read  error message: %v",err)
	}
}

func init() {
	InitJwtKey()
}

func main() {
	Logger,err := InitLogger()
	if err != nil {
		 log.Fatalf("init logger failed,error message %v\n",err)
		return
	}


	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		Logger.Fatal("listen port 8081 failed,error message ",zap.Error(err))
	}


	//添加服务依赖的mongo
	ctx := context.Background()
	var timer   time.Duration = 5*time.Second
	var MaxPoolSize   uint64 = 100
	var MinPoolSize   uint64 = 50
	pref, err := readpref.New(readpref.PrimaryMode)
	if err != nil {
		Logger.Fatal("readpref primary mode,error message ",zap.Error(err))
	}
	con, err := mongo.Connect(ctx, &options.ClientOptions{
		ConnectTimeout:           &timer,
		Hosts:                    []string{
			"47.93.20.75:27017",
		},
		MaxPoolSize:              &MaxPoolSize,
		MinPoolSize:              &MinPoolSize,
		ReadPreference:  pref,
	})
	if err != nil {
		Logger.Fatal("cannot Connect",zap.Error(err))
	}

	db := con.Database("coolcar")
	mon := dao.NewMongo(db)


	//JWT
	pem, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		Logger.Fatal("parse rsa private key from PEM error ",zap.Error(err))
	}
	jwtToken := token.NewJWTTokenGenerate("coolcar/auth",pem)


	srv := grpc.NewServer()
	//该Resolver是实现在Wechat.Srvice的指针上，而不是结构体本身，所以在传递给使用者的时候需要传递wecaht.Service的指针，而不是解耦提本身
	//因为本身并未实现OpenIDResolver接口
	authpb.RegisterAuthServiceServer(srv,&auth.Service{
		OpenIDResolver: &wechat.Service{
			AppSecret: "26c95bc945fd0d7ccf7c473ef5a5e7f8",
			AppID: "wx851020fe449a84e2",
		},
		Logger:                         Logger,
		Mongo: mon,
		TokenGenerator: jwtToken,
		TokenExpire: 2 * time.Hour,
	})


	Logger.Sugar().Fatal("Register auth service failed,error message:",srv.Serve(lis))
}

