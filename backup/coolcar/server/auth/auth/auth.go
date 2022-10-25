package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	TokenGenerator TokenGenerator
	Logger *zap.Logger
	Mongo  *dao.Mongo
	authpb.UnimplementedAuthServiceServer
	TokenExpire time.Duration
}

type OpenIDResolver interface {
	Resolver(code string) (token string,err error)
}

//定义接口的时候并没有说是使用什么具体的Token生成方式，由实现者去定义
type TokenGenerator interface {
	GenerateToken(string,time.Duration)(string,error)
}

func (s *Service)Login(c context.Context,req *authpb.LoginRequest) (res *authpb.LoginResponse,err error) {
	s.Logger.Info("received code",zap.String("code",req.Code))
	openId,err := s.OpenIDResolver.Resolver(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get token error")
	}

	accountID,err := s.Mongo.ResolveAccountID(c,openId)
	if err != nil {
		s.Logger.Error("cannot resolve account id",zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get account id error")
	}

	token, err := s.TokenGenerator.GenerateToken(accountID,s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token ",zap.Error(err))
		return nil, status.Errorf(codes.Internal, "generate token  error")
	}
	//status.Errorf(codes.Unimplemented, "method Login not implemented")
	return &authpb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}

