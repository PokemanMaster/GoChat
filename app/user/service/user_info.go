package service

import (
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
)

type UserInfoService struct {
}

func (service *UserInfoService) UserInfo(id string) resp.Response {
	var user model.User
	db.DB.Model(&user).Where("id = ?", id).First(&user)
	return resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   user,
	}
}
