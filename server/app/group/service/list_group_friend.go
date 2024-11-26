package service

import (
	"github.com/PokemanMaster/GoChat/server/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/server/common/db"
	"github.com/PokemanMaster/GoChat/server/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/server/resp"
)

type GroupFriendListsService struct {
	TargetID uint
}

// List 获取群聊里的好友列表
func (service *GroupFriendListsService) List() resp.Response {
	TargetID := service.TargetID

	var contact []model.Contact
	db.DB.Where("target_id=? and type=2", TargetID).Find(&contact)

	code := e.SUCCESS
	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   contact,
	}
}
