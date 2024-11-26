package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
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
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	return resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   user,
	}
}
