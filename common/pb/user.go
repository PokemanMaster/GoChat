package pb

import (
	"IMProject/app/user/service"
	"IMProject/config"
	pb "IMProject/pb/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func UserPB() {
	grpcServer := grpc.NewServer()                        // 创建 gRPC 服务器
	userService := service.GetUserSrv()                   // 注册 UserSrv
	pb.RegisterUserServiceServer(grpcServer, userService) // 注册中心
	grpcAddress := config.Conf.Services["user"].Addr[0]   // 设置监听地址
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil { // 启动 gRPC 服务器
		logrus.Fatalf("failed to serve: %v", err)
	}
}
