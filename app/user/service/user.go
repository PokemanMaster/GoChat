package service

import (
	"IMProject/app/user/dao"
	pb "IMProject/pb/user"
	"IMProject/pkg/e"
	"context"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
	pb.UnimplementedUserServiceServer
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) Register(ctx context.Context, req *pb.UserRequest) (res *pb.UserCommonResponse, err error) {
	res = new(pb.UserCommonResponse)
	res.Code = e.SUCCESS
	err = dao.NewUserDao(ctx).CreateUser(req)
	if err != nil {
		res.Code = e.ERROR
		return
	}
	res.Data = e.GetMsg(int(res.Code))
	return
}

func (u *UserSrv) Login(ctx context.Context, req *pb.UserRequest) (res *pb.UserDetailResponse, err error) {
	res = new(pb.UserDetailResponse)
	res.Code = e.SUCCESS
	user, err := dao.NewUserDao(ctx).UserLogin(req)
	if err != nil {
		res.Code = e.ERROR
		res.Msg = err.Error() // 将错误信息返回给客户端
		return
	}
	res.UserDetail = &pb.UserResponse{
		UserId:   int64(user.ID),
		UserName: user.UserName,
		Identity: user.Identity,
	}
	res.Msg = e.GetMsg(int(res.Code))
	return res, nil
}

func (u *UserSrv) Update(ctx context.Context, req *pb.UserRequest) (res *pb.UserDetailResponse, err error) {
	res = new(pb.UserDetailResponse)
	res.Code = e.SUCCESS
	user, err := dao.NewUserDao(ctx).UpdateUser(req)
	if err != nil {
		res.Code = e.ERROR
		res.Msg = err.Error()
		return
	}
	res.UserDetail = &pb.UserResponse{
		UserId:   int64(user.ID),
		UserName: user.UserName,
		Identity: user.Identity,
	}
	res.Msg = e.GetMsg(int(res.Code))
	return res, nil
}
