package main

import (
	"crypto/tls"
	"go_grpc_gorm_micro/api"
	"go_grpc_gorm_micro/lib/gateway"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go_grpc_gorm_micro/lib/global"
	"go_grpc_gorm_micro/lib/middleware/auth"
	"go_grpc_gorm_micro/lib/middleware/cred"
	"go_grpc_gorm_micro/lib/middleware/recovery"
	"go_grpc_gorm_micro/lib/middleware/zap"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go_grpc_gorm_micro/proto/proto"
	"go_grpc_gorm_micro/initialize"
)

func main() {

	// 监听本地端口
	listener, err := net.Listen(global.CURD_CONFIG.System.Network, global.CURD_CONFIG.System.Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)

	initialize.Gorm()     // gorm连接数据库

	//grpcServer := grpc.NewServer()

	// 在gRPC服务器注册我们的服务
	pb.RegisterSysApisServiceServer(grpcServer, &api.SysApis{})

	log.Println(global.CURD_CONFIG.System.Address + " net.Listing whth TLS and token...")

	//使用gateway把grpcServer转成httpServer
	httpServer := gateway.ProvideHTTP(global.CURD_CONFIG.System.Address, grpcServer)

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	if err = httpServer.Serve(tls.NewListener(listener, httpServer.TLSConfig)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
