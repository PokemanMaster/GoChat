package service

import (
	"github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
)

type GroupFriendListsService struct {
	TargetID uint
}

// List 获取群聊里的好友列表
func (service *GroupFriendListsService) List() resp.Response {
	TargetID := service.TargetID

	var contact []model.Contact
	db.DB.Where("target_id=? and type=2", TargetID).Find(&contact)

	code := e2.SUCCESS
	return resp.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   contact,
	}
}
