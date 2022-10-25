package server

import (
	"coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type RunGrpcServerConfig struct {
	NetWork string
	Addr string
	Logger *zap.Logger
	Name string
	PubKey  string
	RegisterFunc func(server *grpc.Server) //这里的思想值得关注
}

func RUNGrpcServer(c *RunGrpcServerConfig)  (err error) {

	nameField := zap.String("name",c.Name)

	//统一注册服务
	lis, err := net.Listen(c.NetWork, c.Addr)
	if err != nil {
		c.Logger.Fatal("can not listen",nameField,zap.Error(err))
		return
	}


	//由于每个服务之间都需要验证token，所以需要在各个服务启动之前进行验证
	//在grpc中，可以在启动grpc服务的时候注入拦截器，使用拦截器的方式进行类似中间件的注入一样
	//来在真正到达调用服务之前可以进行一些必要的前置操作，如验证token、参数规则校验等
	//grpc中的拦截器，是在NewServcer中传递参数,这里使用简单的拦截器ChainUnaryInterceptor
	//他的参数是type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
	//也就是说只要传递了一个func和上边一样就可以，由于拦截器属于公用工具，所以最好的方式就是写在公共位置shared中

	//因为登录不需要做什么验证，所以需要特出处理类似的服务，将拦截器过滤掉
	var ops []grpc.ServerOption
	if c.PubKey != "" {
		info, err := auth.InterceptorInfo(c.PubKey)
		if err != nil {
			c.Logger.Fatal("can not InterceptorInfo",nameField,zap.Error(err))
			return
		}
		interceptor := grpc.ChainUnaryInterceptor(info)
		ops = append(ops,interceptor)
	}

	//创建grpc服务
	srv := grpc.NewServer(ops...)

	//调用外部传递的函数
	c.RegisterFunc(srv)//这里的思想值得关注

	/*********************不能做到更好的完全统一处理，因为服务不同，注册服务的函数不通，所以就索性都调用外部的函数，交由外部处理**************************/
	//tripSrv := &trip.Service{
	//	Logger:                         c.Logger,
	//}

	//这个由于各个服务都不太确定,但是所有的服务注册函数都是一样的函数类型
	//即这种情况下就可以进行暴露给外部函数，去让外部决定操作
	//所以就传递给外部一个函数，将内部的信息告诉外部，然后让外部去操作
	//rentalpb.RegisterTripServiceServer(srv,tripSrv)
	/*********************不能做到更好的完全统一处理，因为服务不同，注册服务的函数不通，所以就索性都调用外部的函数，交由外部处理**************************/

	return srv.Serve(lis)
}
