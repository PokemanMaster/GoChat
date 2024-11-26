package service

import (
	"github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

type UserInfoService struct {
}

func (service *UserInfoService) UserInfo(id string) resp.Response {
	var user model.User

	err := db.DB.Model(&user).Where("id = ?", id).First(&user).Error
	if err != nil {
		zap.L().Error("获取用户信息失败", zap.String("app.user.service.user_info", err.Error()))
		return resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	return resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
		Data:   user,
	}
}
