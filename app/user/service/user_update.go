package service

import (
	"github.com/PokemanMaster/GoChat/server/app/user/model"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"github.com/asaskevich/govalidator"
	"go.uber.org/zap"
)

// UserUpdateService 前端请求过来的数据
type UserUpdateService struct {
	ID        uint
	UserName  string
	Password  string
	Telephone string
	Avatar    string
	Email     string
}

func (service *UserUpdateService) UserUpdate() *resp.Response {
	user := model.User{}
	user.ID = service.ID
	user.UserName = service.UserName
	user.Password = service.Password
	user.Telephone = service.Telephone
	user.Avatar = service.Avatar
	user.Email = service.Email

	// 更新用户
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		zap.L().Error("更新用户失败", zap.String("app.user.service.user_update", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_MATCHED_USERNAME,
			Msg:    e2.GetMsg(e2.ERROR_MATCHED_USERNAME),
		}
	}

	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
	}
}
