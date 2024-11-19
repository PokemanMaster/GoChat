package service

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/resp"
)

type UserInfoService struct {
}

func (service *UserInfoService) UserInfo(id string) resp.Response {
	var user model.UserBasic
	db.DB.Model(&user).Where("id = ?", id).First(&user)
	return resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   user,
	}
}
